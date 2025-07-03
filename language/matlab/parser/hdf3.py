import h5py
import numpy as np
from datetime import datetime, timedelta
from typing import Dict
from utils.log import info, warn, error
import csv

class HDF5MatLoader(object):
    def load(self, file_path: str) -> Dict:
        with h5py.File(file_path, 'r') as f:
            result = {}
            for key in f.keys():
                node = f[key]
                if isinstance(node, h5py.Group) and 'MATLAB_class' in node.attrs and node.attrs[
                    'MATLAB_class'] == b'struct':
                    result[key] = self._extract_struct_array(node, f)
                else:
                    result[key] = self._resolve_node(node, f)
            return result

    def _extract_struct_array(self, group: h5py.Group, file):
        """提取 HDF5 结构体数组，转换为 List[Dict]"""
        num_entries = next(iter(group.values())).shape[0]
        records = []

        for i in range(num_entries):
            entry = {}
            for field_name, dataset in group.items():
                try:
                    ref = dataset[i].item() if isinstance(dataset[i], np.ndarray) else dataset[i]
                    if isinstance(ref, h5py.Reference):
                        val = self._resolve_node(file[ref], file, field_name)
                    else:
                        val = ref
                    entry[field_name] = val
                except Exception as e:
                    error(f"extract struct fields failed: {e}")
                    entry[field_name] = None
            records.append(entry)
        return records

    def _resolve_node(self, node, file, field_name=''):
        try:
            if isinstance(node, h5py.Dataset):
                val = node[()]

                # uint16 字符类型处理
                if isinstance(val, np.ndarray) and val.dtype == np.uint16:
                    try:
                        return ''.join(chr(c) for c in val.flatten() if c > 0)
                    except Exception as e:
                        error(f"resolve node fields failed, err: {e}")
                        return val.tolist()

                # 引用数组处理
                if isinstance(val, np.ndarray) and val.dtype == object:
                    result = []
                    for ref in val:
                        if isinstance(ref, h5py.Reference):
                            target = file[ref]
                            target_val = target[()]

                            if self.__is_mcos_encoded(target_val):
                                result.append(f"MCOS: {field_name}")
                            elif isinstance(target_val, np.ndarray) and target_val.shape == (1, 1):
                                result.append(target_val.item())
                            elif isinstance(target_val, np.ndarray) and target_val.shape == (1,):
                                result.append(target_val[0])
                            else:
                                result.append(target_val.tolist())
                        else:
                            result.append(ref)
                    return result

                # 标量或常规数组
                if isinstance(val, np.ndarray):
                    if self.__is_mcos_encoded(val):
                        return f"MCOS: {field_name}"
                    if val.shape == (1, 1):
                        return val.item()
                    return val.tolist()

            elif isinstance(node, h5py.Group):
                return {k: self._resolve_node(v, file) for k, v in node.items()}

            return str(node)
        except Exception as e:
            error(f"resolve node fields failed, err: {e}")
            return None

    @staticmethod
    def __is_mcos_encoded(val: np.ndarray) -> bool:
        return (
                isinstance(val, np.ndarray) and
                val.dtype == np.uint32 and
                val.shape[0] == 1 and
                val.shape[1] >= 6 and
                val[0, 0] > 3_000_000_000
        )


def convert_timestamp(value: float) -> str:
    base_time = datetime(2011, 1, 1, 1, 1, 1)
    timestamp = base_time + timedelta(seconds=value)
    return timestamp.strftime("%Y-%m-%d %H:%M:%S.") + str(timestamp.microsecond)[:3]


if __name__ == '__main__':
    filepath = "../practise/testcase/mock_data.mat"
    output = "data.csv"
    data = HDF5MatLoader().load(filepath)
    print(type(data))
    keys = list(data.keys())

    if len(keys) == 0:
        raise Exception("数据为空")

    handle_key = keys[0]
    columns = len(data[handle_key])
    rows = len(data[handle_key][0])
    print(f"数据行数：{rows}，列数：{columns}")

    header = [
        "id",
        "tm",
        "f1",
        "f2",
        "f3",
        "f4",
        "f5",
        "f6",
        "f7",
        "f8",
        "f9",
        "f10",
        "f11",
        "f12",
        "f13",
        "f14",
        "f15",
        "updated_at"
    ]

    start_time = datetime.now()
    with open(output, 'w', newline='', encoding='utf-8') as f:
        writer = csv.writer(f)
        writer.writerow(header)

        for i in range(rows):
            current_time = datetime.now()
            new_time = current_time + timedelta(seconds=0.01)

            row = [
                1,
                convert_timestamp(data[handle_key][6][i]),
                data[handle_key][0][i],
                data[handle_key][1][i],
                data[handle_key][2][i],
                data[handle_key][3][i],
                data[handle_key][4][i],
                data[handle_key][5][i],
                data[handle_key][8][i],
                data[handle_key][9][i],
                data[handle_key][10][i],
                data[handle_key][11][i],
                data[handle_key][12][i],
                data[handle_key][13][i],
                data[handle_key][14][i],
                data[handle_key][15][i],
                data[handle_key][16][i],
                new_time.strftime("%Y-%m-%d %H:%M:%S.") + str(new_time.microsecond)[:3]
            ]
            writer.writerow(row)
            print(f"行 {i + 1}: {row}")
            print("-" * 40)  # 每行输出分隔线

    end_time = datetime.now()

    total_time = (end_time - start_time).total_seconds()
    print(f"数据解析完成，耗时：{total_time} 秒")
