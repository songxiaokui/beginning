from typing import Any, List, Dict
from datetime import datetime
import numpy as np
import h5py
import scipy.io.matlab._mio5_params as mio5
from .log import info, warn, error


def is_timestamp_column(val: Any) -> bool:
    try:
        ts = float(val)
        # 合理的 Unix 时间戳范围 (2000-01-01 到 2100-01-01)
        return 946684800 <= ts <= 4102444800
    except:
        return False


def normalize_records(raw: Any) -> List[Dict]:
    if not isinstance(raw, (list, tuple, np.ndarray)):
        raw = [raw]

    records = [flatten_mat_struct(r) for r in raw]

    if records and isinstance(records[0], list):
        records = [item for sublist in records for item in sublist]

    for row in records:
        if not hasattr(row, "items"):
            return records
        for k, v in row.items():
            if is_timestamp_column(v):
                try:
                    row[k] = datetime.utcfromtimestamp(float(v)).strftime("%Y-%m-%d %H:%M:%S.%f")[:-3]
                except:
                    pass
    return records


def flatten_mat_struct(mat_struct: Any) -> Any:
    """递归解析结构体，包括 MatlabOpaque 类型（MCOS 对象）"""

    if isinstance(mat_struct, (list, tuple, np.ndarray)):
        return [flatten_mat_struct(item) for item in mat_struct]

    if isinstance(mat_struct, mio5.MatlabOpaque):
        try:
            class_name = mat_struct[2].decode() if isinstance(mat_struct[2], bytes) else str(mat_struct[2])
        except Exception as e:
            class_name = 'Unknown'
            warn(f"failed to decode MatlabOpaque class name: {e}")
        return f"MCOS: {class_name}"

    if isinstance(mat_struct, dict):
        result = {}
        for k, v in mat_struct.items():
            if isinstance(v, mio5.MatlabOpaque):
                try:
                    cls = v[2].decode() if isinstance(v[2], bytes) else str(v[2])
                except Exception:
                    cls = 'Unknown'
                warn(f"field '{k}' contains unrecognized MCOS type: {cls}")
                result[k] = f"MCOS: {k}"
            elif isinstance(v, (dict, list, tuple, np.ndarray)):
                result[k] = flatten_mat_struct(v)
            else:
                result[k] = v
        return result

    return mat_struct


def is_hdf5(file_path: str) -> bool:
    try:
        with h5py.File(file_path, 'r'):
            return True
    except Exception as e:
        error(f"failed to open file as HDF5: {e}")
        return False
