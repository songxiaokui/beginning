from OCC.Extend.TopologyUtils import TopologyExplorer
from OCC.Core.STEPControl import STEPControl_Reader
from OCC.Core.BRep import BRep_Tool
from OCC.Core.TopoDS import topods


def read_step_file(filename):
    """读取STEP文件并返回一个TopoDS_Shape对象"""
    step_reader = STEPControl_Reader()
    step_reader.ReadFile(filename)
    step_reader.TransferRoot()
    shape = step_reader.Shape()
    return shape


def explore_shape(shape):
    """探索并打印形状中的所有面和体的信息"""
    explorer = TopologyExplorer(shape)

    print("Solid total count: ", explorer.number_of_solids())
    for solid_index, solid in enumerate(explorer.solids(), start=1):
        print(f"Solid {solid_index} found")
        for face_index, face in enumerate(explorer.faces_from_solids(solid),
                                          start=1):
            print(f"  Face {solid_index}.{face_index} found")
            for edge_index, edge in enumerate(explorer.edges_from_face(face),
                                              start=1):
                print(
                    f"    Edge {solid_index}.{face_index}.{edge_index} found")
                for vertex_index, vertex in enumerate(
                        explorer.vertices_from_edge(edge), start=1):
                    pnt = BRep_Tool.Pnt(topods.Vertex(vertex))
                    print(
                        f"      Vertex {solid_index}.{face_index}.{edge_index}.{vertex_index} found at ({pnt.X()}, {pnt.Y()}, {pnt.Z()})"
                    )


# 替换为你的STEP文件路径
if __name__ == '__main__':
    filename = "Box_2.step"
    # filename = "PanShape2.step"
    shape = read_step_file(filename)
    if shape:
        explore_shape(shape)
