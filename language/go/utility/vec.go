package main

import (
	"encoding/json"
	"fmt"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
	"math"
	"os"
	"time"
)

type Location struct {
	X float64
	Y float64
	Z float64
}

func NewLocation(x, y, z float64) Location {
	return Location{
		X: x,
		Y: y,
		Z: z,
	}
}

type FieldData struct {
	Coordinates []float64              `json:"coordinates"` // 坐标
	Fields      map[string]interface{} `json:"fields"`      // 场数据
}

func (fd FieldData) WarpCoordinate() Location {
	return Location{
		X: fd.Coordinates[0],
		Y: fd.Coordinates[1],
		Z: fd.Coordinates[2],
	}
}

// search field data by coordinates

type SearchFielder interface {
	Load() error
	Search(location Location) (any, error)
}

// EuclideanDistancesImpl traverse all data using Euclidean distances
type EuclideanDistancesImpl struct {
	InputFile string
	data      []FieldData
}

func NewEuclideanDistancesImpl(input string) SearchFielder {
	return &EuclideanDistancesImpl{
		InputFile: input,
		data:      make([]FieldData, 0),
	}
}

func (ed *EuclideanDistancesImpl) Load() error {
	startTime := time.Now()
	defer func() {
		fmt.Printf("load file consume: %.2f s", time.Now().Sub(startTime).Seconds())
	}()
	dataByte, err := os.ReadFile(ed.InputFile)
	fmt.Printf("read file consume: %.2f s\n", time.Now().Sub(startTime).Seconds())
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, &ed.data)
}

func (ed *EuclideanDistancesImpl) Search(location Location) (any, error) {
	var minDistances float64 = math.Inf(1)
	var Data interface{} = nil
	startTime := time.Now()
	for _, source := range ed.data {
		d := ed.distance(source.WarpCoordinate(), location)
		if d < minDistances {
			Data = source
			minDistances = d
		}
	}
	endTime := time.Now()
	fmt.Printf("search consume: %.2f s", endTime.Sub(startTime).Seconds())
	return Data, nil
}

func (ed *EuclideanDistancesImpl) distance(source, target Location) float64 {
	return math.Sqrt(math.Pow(source.X-target.X, 2) + math.Pow(source.Y-target.Y, 2) + math.Pow(source.Z-target.Z, 2))
}

// KDTreeImpl Implementation using KD Tree
type KDTreeImpl struct {
	InputFile string
	data      []FieldData
	tree      *kdtree.KDTree
}

// Point3D specifies one element of the k-d tree.
type Point3D struct {
	FieldData
}

// Dimensions returns the total number of dimensions
func (p3 *Point3D) Dimensions() int {
	return len(p3.Coordinates)
}

// Dimension returns the value of the i-th dimension
func (p3 *Point3D) Dimension(i int) float64 {
	return p3.Coordinates[i]
}

func NewPoint(fd FieldData) kdtree.Point {
	return &Point3D{
		FieldData: fd,
	}
}

func NewKDTreeImpl(input string) SearchFielder {
	p1 := make([]kdtree.Point, 0)
	return &KDTreeImpl{
		InputFile: input,
		data:      make([]FieldData, 0),
		tree:      kdtree.New(p1),
	}
}

func (kd *KDTreeImpl) Load() error {
	startTime := time.Now()
	defer func() {
		fmt.Printf("load file consume: %.2f s", time.Now().Sub(startTime).Seconds())
	}()
	dataByte, err := os.ReadFile(kd.InputFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataByte, &kd.data)
	if err != nil {
		return err
	}
	fmt.Printf("%.2f\n", time.Now().Sub(startTime).Seconds())
	// storage
	for _, v := range kd.data {
		kd.tree.Insert(NewPoint(v))
	}
	return nil
}

func (kd *KDTreeImpl) Search(location Location) (any, error) {
	startTime := time.Now()
	result := kd.tree.KNN(&points.Point{Coordinates: []float64{location.X, location.Y, location.Z}}, 1)
	endTime := time.Now()
	fmt.Printf("search consume: %.2f s", endTime.Sub(startTime).Seconds())
	return result[0], nil
}
