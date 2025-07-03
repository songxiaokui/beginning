import argparse
from loader.scipy_loader import ScipyMatLoader
from loader.hdf5_loader import HDF5MatLoader
from converter.export_csv import export_to_csv
from utils.format import is_hdf5, normalize_records
from utils.log import warn, info


def main():
    parser = argparse.ArgumentParser(description="matlab .mat adapter")
    parser.add_argument('--input', required=True, help='path to .mat file')
    parser.add_argument('--output', required=True, help='output CSV path')
    parser.add_argument('--field', required=True, help='output CSV path')
    args = parser.parse_args()

    loader = HDF5MatLoader() if is_hdf5(args.input) else ScipyMatLoader()
    data = loader.load(args.input)

    if args.field in data:
        records = normalize_records(data[args.field])
        headers = list(records[0].keys()) if records and hasattr(records[0], 'keys') else []
        export_to_csv(records, args.output, headers=headers)
        info(f"exported {args.field} to {args.output} with headers: {headers}")
    else:
        warn(f"{args.field} field not found. Attempting to export other fields...")
        for key in data.keys():
            records = normalize_records(data[key])
            headers = list(records[0].keys()) if records else []
            output_path = f"{args.output}_{key}.csv"
            export_to_csv(records, output_path, headers=headers)
            info(f"exported '{key}' to {output_path} with headers: {headers}")


if __name__ == '__main__':
    main()
