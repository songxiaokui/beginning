import h5py
import numpy as np
import pandas as pd
from typing import Dict, List, Any
from utils.log import error, info


def decode_uint16_strings(arr: np.ndarray) -> List[str]:
    """将 uint16 编码的二维数组解码为字符串列表"""
    return [''.join(chr(c) for c in row if c > 0) for row in arr]


def resolve_reference(f: h5py.File, ref: h5py.Reference, field: str) -> Any:
    """解析 MCOS 引用或标量值"""
    try:
        node = f[ref]
        val = node[()]

        if isinstance(val, np.ndarray):
            if val.dtype == np.uint32 and val.shape in [(1, 6), (1, 7)]:
                return f"MCOS: {field}"
            elif val.dtype == np.float64 and val.shape == (1, 1):
                return val.item()
            elif val.shape == (1,):
                return val[0]
            elif val.dtype == np.uint16:
                return ''.join(chr(c) for c in val.flatten() if c > 0)
            else:
                return val.tolist()
        return val
    except Exception as e:
        return f"[ERR: {e}]"


def load_mat_v7_3_chunk_data(file_path: str) -> pd.DataFrame:
    """加载 .mat (v7.3, HDF5) 文件中的 chunk_data 字段并转换为 DataFrame"""
    with h5py.File(file_path, 'r') as f:
        if "chunk_data" not in f:
            raise ValueError("File does not contain 'chunk_data' group.")

        chunk_data = f["chunk_data"]
        fields = list(chunk_data.keys())

        result: Dict[str, List[Any]] = {}

        for field in fields:
            dataset = chunk_data[field]
            values = dataset[()]
            resolved_values: List[Any] = []

            try:
                if dataset.dtype == np.uint16:
                    resolved_values = decode_uint16_strings(values)
                elif dataset.dtype == object and values.ndim == 2 and isinstance(values[0][0], h5py.Reference):
                    resolved_values = [
                        resolve_reference(f, ref_row[0], field) for ref_row in values
                    ]
                else:
                    resolved_values = values.tolist()
            except Exception as e:
                error(f"failed to process field '{field}': {e}")
                resolved_values = [None] * dataset.shape[0]

            result[field] = resolved_values

        df = pd.DataFrame(result)
        df = df.reindex(columns=fields)
        return df


def main():
    file_path = "../../../../scripts/testcase/output_v7.3/chunk_1_v7.3_1.mat"
    output_path = "../../abc.csv"

    try:
        df = load_mat_v7_3_chunk_data(file_path)
        df.to_csv(output_path, index=False)
        info(f"Data saved to {output_path}")
    except Exception as e:
        error(f"Failed to process file: {e}")


if __name__ == '__main__':
    main()
