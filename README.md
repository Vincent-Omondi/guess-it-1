<a name="readme-top"></a>


# Math Skills
## Statistical Analysis Tool

This project is a command-line tool written in Go that performs statistical analysis on numerical data read from a file. It calculates the average, median, variance, and standard deviation of the data and provides the results in a user-friendly format.

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## About The Project
The primary objective of this project is to create a robust and efficient statistical analysis tool that can handle large datasets with ease. The tool is designed to read numerical data from a file, perform various statistical calculations, and display the results in a readable format.

### Built With
<img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="60" height="60">




<p align="right">(<a href="#math-skills">back to top</a>)</p>


## Getting Started

To get started with Math Skills, follow these instructions:

### Prerequisites
- Go installed on your machine.
- Basic understanding of Go programming language.


### Installation
Clone the repository:
```sh
git clone https://github.com/Vincent-Omondi/math-skills.git
``` 

Navigate to the project directory:

```sh
cd math-skills
```
<p align="right">(<a href="#math-skills">back to top</a>)</p>


## Usage

To use Math Skills, follow these steps:

1. Prepare a text file containing the dataset you want to analyze. Each line of the file should represent one value of the dataset.

2. Compile the program using the following command

```sh
go build
```
3. Run the compiled program with the file path as an argument

```sh
./math-skills <file_path>
```
Replace <file_path> with the path to the file containing the numerical data you want to analyze. The file should have one number per line, and non-numeric values and overflow will be ignored.
Example:
```sh
./math-skills data.txt
```

The program will output the statistics in the following manner (the following numbers are only examples)

```
User$ ./math-skills data.txt
Average: 147
Median: 146
Variance: 769
Standard Deviation: 28
```
All values are rounded to the nearest integer.
The program accepts a wide range of files including the following:

- **Plain Text Formats**: `.txt`, `.csv`
- **Spreadsheet Formats**: `.xls`, `.xlsx`, `.ods`
- **Database Formats**: `.sqlite`, `.dbf`
- **JSON and XML Formats**: `.json`, `.xml`
- **Specialized Data Formats**: `.hdf`, `.nc`
- **Statistical Software Formats**: `.sav`, `.dta`
- **Compressed Formats**: `.zip`, `.gz`
- **Web and API Formats**: `.html`, `.json` (from API calls)
- **Geographic Information System (GIS) Formats**: `.shp`, `.kml`

<p align="right">(<a href="#math-skills">back to top</a>)</p>

## RoadMap
* Add support for different input file formats (e.g., CSV, xlsx)
* Implement additional statistical measures (e.g., quartiles, percentiles)
* Enhance error handling and input validation
* Optimization for large datasets.
* Develop a graphical user interface (GUI) for better user experience

<p align="right">(<a href="#math-skills">back to top</a>)</p>


## Contributing

Contributions are welcome! If you want to contribute to Math Skills, follow these steps:

1. Fork the project.
2. Create your feature or bugfix branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature/bugfix'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

<p align="right">(<a href="#math-skills">back to top</a>)</p>


## License
This project is licensed under the MIT License.

<p align="right">(<a href="#math-skills">back to top</a>)</p>

## Contact
If you have any questions or need further assistance, please feel free to contact the project maintainer:

[X](https://tweeter.com/vinomondi_1)

[Github](https://github.com/Vincent-Omondi/)

<p align="right">(<a href="#math-skills">back to top</a>)</p>


## Acknowledgments

Special thanks to the Zone01 for their valuable resources and support.

<p align="right">(<a href="#math-skills">back to top</a>)</p>
