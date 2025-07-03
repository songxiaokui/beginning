import scipy.io
from typing import Dict
from .base_loader import MatLoader


class ScipyMatLoader(MatLoader):
    def load(self, file_path: str) -> Dict:
        return scipy.io.loadmat(file_path, simplify_cells=True)
