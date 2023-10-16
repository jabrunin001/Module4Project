# Module4Project

California Housing Prices Analysis
This project analyzes California housing prices data, deriving summary statistics for seven key variables. The analysis is performed using three different scripts: R, Python (with Pandas), and Go.

Overview
The data for this project is sourced from a study on California housing prices (Miller, 2015). The dataset contains the following columns:

value: The house value.
income: Income level.
age: Age of the house.
rooms: Number of rooms.
bedrooms: Number of bedrooms.
pop: Population in the area.
hh: Household size.
The aim is to compute three summary statistics for each of these columns:

Minimum
Maximum
Mean
Scripts
1. R Script (runHouses.R)
This script reads the data from housesInput.csv, calculates summary statistics using R's built-in functions, and writes the results to housesOutputR.txt.

Running the script:

bash
Copy code
Rscript runHouses.R
2. Python Script (runHouses.py)
This script uses the Pandas library to process the data. It reads from housesInput.csv, computes the summary statistics, and saves the results in housesOutputPy.txt.

Running the script:

bash
Copy code
python3 runHouses.py
3. Go Command-Line Application
The Go application performs the same analysis as the R and Python scripts. It reads the data, computes the summary statistics, and writes the results to housesOutputGo.txt.

Running the application:
(assuming the compiled executable is named housesGoApp)

bash
Copy code
./housesGoApp
Benchmarking
To compare the performance of the three scripts, each script (or application) is run 100 times in succession. The total CPU processing time is then recorded for each.

Timing on Unix-like systems:

bash
Copy code
time Rscript runHouses.R
time python3 runHouses.py
time ./housesGoApp
Timing on Windows (using PowerShell):

powershell
Copy code
Measure-Command { Rscript.exe runHouses.R }
Measure-Command { python runHouses.py }
Measure-Command { .\housesGoApp.exe }
Results
After running the scripts, the results can be compared by examining the output files (housesOutputR.txt, housesOutputPy.txt, and housesOutputGo.txt) to ensure consistency in the computed statistics.

You can save the above content into a README.md file and include it in the root directory of your project. Adjustments can be made based on any additional details or specifics you'd like to include.
