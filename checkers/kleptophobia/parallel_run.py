from concurrent.futures import ProcessPoolExecutor
from subprocess import check_output


def run_checker():
    with ProcessPoolExecutor(32) as executor:
        for _ in range(100):
            executor.submit(check_output, ["python3", "checker.py", "test", "localhost"])
    executor.shutdown(wait=True)


def main():
    run_checker()


if __name__ == '__main__':
    main()
