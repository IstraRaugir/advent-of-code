from pathlib import Path

def read_raw(file_name: str) -> list[str]:
    path = Path(__file__).parent / file_name
    with path.open("r") as input:
        return input.readlines()

def read_int(file_name: str) -> list[int]:
    return list(map(int, read_raw(file_name)))
