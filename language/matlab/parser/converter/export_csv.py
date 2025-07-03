import pandas as pd
from typing import List, Dict


def export_to_csv(data: List[Dict], path: str, headers: List[str] = None):
    """导出数据到 CSV"""
    df = pd.DataFrame(data)
    if headers:
        df = df.reindex(columns=headers)
    df.to_csv(path, index=False)
