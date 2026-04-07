import argparse
import os
import shutil


def create_path(sequence):
    seq_num = int(sequence)
    root_dir = "042-100"
    sub_dir = "042-050"

    if seq_num > 100:
        root_num = (seq_num // 100 + 1) * 100
        root_dir = f"{root_num-99:03d}-{root_num:03d}"

    if seq_num > 50:
        sub_num = (seq_num // 10  * 10) + 1
        sub_dir = f"{sub_num:03d}-{sub_num+9:03d}"

    return root_dir, sub_dir


def main():
    # コマンドライン引数の処理
    parser = argparse.ArgumentParser()
    parser.add_argument("category", help="category of the problem")
    parser.add_argument("sequence", help="sequence number of the problem")
    parser.add_argument("level", help="difficulty level of the problem")

    args = parser.parse_args()
    category = args.category.upper()
    sequence = args.sequence
    level = args.level.upper()
    # コピー元とコピー先のパスを作成
    root_dir, sub_dir = create_path(sequence)
    source_dir = f"..\\{category}\\{root_dir}\\{sub_dir}"
    source_file = f"{source_dir}\\{category}-{sequence}-{level}.go"

    target_file = "main.go"
    # コピー元のディレクトリとファイルが存在する場合はコピー
    if os.path.exists(source_dir) and os.path.exists(source_file):
        shutil.copy2(source_file, target_file)
    else:
        print(f"Error: Source file '{source_file}' does not exist.")


if __name__ == "__main__":
    main()
