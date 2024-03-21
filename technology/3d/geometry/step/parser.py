from OCC.Extend.DataExchange import read_step_file
from OCC.Core.TopoDS import topods_Vertex, topods_Edge, topods_Face, topods_Solid
from OCC.Core.TopExp import TopExp_Explorer
from OCC.Core.TopAbs import TopAbs_VERTEX, TopAbs_EDGE, TopAbs_FACE, TopAbs_SOLID
from OCC.Core.BRep import BRep_Tool
import os


class StepGeomProcessor:
    def __init__(self, filepath):
        self.shape = read_step_file(filepath)
        self.point_id_map = {}
        self.line_id_map = {}
        self.face_id_map = {}
        self.solid_id_map = {}
        self.id_generator = self.id_gen(self)
        self.point_id_location = {}
        self.line_id_location = {}
        self.face_id_location = {}
        self.solid_id_location = {}

    def __convert(self):
        for k, v in self.point_id_map.items():
            self.point_id_location[v] = k
        for k, v in self.line_id_map.items():
            self.line_id_location[v] = k
        for k, v in self.face_id_map.items():
            self.face_id_location[v] = k
        for k, v in self.solid_id_map.items():
            self.solid_id_location[v] = k

    @staticmethod
    def id_gen(self):
        n = 1
        while True:
            yield n
            n += 1

    def process(self):
        self.process_vertices()
        self.process_edges()
        self.process_faces()
        self.process_solids()
        self.__convert()

    def process_vertices(self):
        explorer = TopExp_Explorer(self.shape, TopAbs_VERTEX)
        while explorer.More():
            vertex = topods_Vertex(explorer.Current())
            point = BRep_Tool.Pnt(vertex)
            point_tuple = (point.X(), point.Y(), point.Z())
            if point_tuple not in self.point_id_map:
                self.point_id_map[point_tuple] = next(self.id_generator)
            explorer.Next()

    def process_edges(self):
        explorer = TopExp_Explorer(self.shape, TopAbs_EDGE)
        while explorer.More():
            edge = topods_Edge(explorer.Current())
            vertices = self.get_vertices(edge)
            edge_tuple = tuple(sorted(vertices))
            if edge_tuple not in self.line_id_map:
                self.line_id_map[edge_tuple] = next(self.id_generator)
            explorer.Next()

    def process_faces(self):
        explorer = TopExp_Explorer(self.shape, TopAbs_FACE)
        while explorer.More():
            face = topods_Face(explorer.Current())
            edges = self.get_edges(face)
            face_tuple = tuple(sorted(edges))
            if face_tuple not in self.face_id_map:
                self.face_id_map[face_tuple] = next(self.id_generator)
            explorer.Next()

    def process_solids(self):
        explorer = TopExp_Explorer(self.shape, TopAbs_SOLID)
        while explorer.More():
            solid = topods_Solid(explorer.Current())
            faces = self.get_faces(solid)
            solid_tuple = tuple(sorted(faces))
            if solid_tuple not in self.solid_id_map:
                self.solid_id_map[solid_tuple] = next(self.id_generator)
            explorer.Next()

    def get_vertices(self, shape):
        vertices = []
        explorer = TopExp_Explorer(shape, TopAbs_VERTEX)
        while explorer.More():
            vertex = topods_Vertex(explorer.Current())
            point = BRep_Tool.Pnt(vertex)
            point_tuple = (point.X(), point.Y(), point.Z())
            vertices.append(self.point_id_map[point_tuple])
            explorer.Next()
        return vertices

    def get_edges(self, shape):
        edges = []
        explorer = TopExp_Explorer(shape, TopAbs_EDGE)
        while explorer.More():
            edge = topods_Edge(explorer.Current())
            vertices = self.get_vertices(edge)
            edge_tuple = tuple(sorted(vertices))
            edges.append(self.line_id_map[edge_tuple])
            explorer.Next()
        return edges

    def get_faces(self, shape):
        faces = []
        explorer = TopExp_Explorer(shape, TopAbs_FACE)
        while explorer.More():
            face = topods_Face(explorer.Current())
            edges = self.get_edges(face)
            face_tuple = tuple(sorted(edges))
            faces.append(self.face_id_map[face_tuple])
            explorer.Next()
        return faces


if __name__ == '__main__':
    current_path = os.path.dirname(os.path.abspath(__file__))
    # geom_file_path = os.path.join(current_path, "PanShape2.step")
    geom_file_path = os.path.join(current_path, "Box_2.step")
    print("文件路径: ", geom_file_path)
    step = StepGeomProcessor(geom_file_path)
    step.process()
    print(step.point_id_map)
    print(step.point_id_location)
    print(step.line_id_map)
    print(step.line_id_location)
    print(step.face_id_map)
    print(step.face_id_location)
    print(step.solid_id_map)
    print(step.solid_id_location)