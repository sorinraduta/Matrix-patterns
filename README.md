<p align="center">
  <img height="100" width="100" src="matrix.svg" />
  <h3 align="center">Matrix patterns</h3>
</p>

## Instructions
Uncomment a formula from `calculateFormula` function and run `go run main.go`;

## Formulas
<img height="200" width="200" src="screenshots/1.png" /><br />
maxColumns * (currentRow-1) + currentColumn

---
<img height="200" width="200" src="screenshots/2.png" /><br />
maxColumns * (currentRow - 1) + maxColumns - currentColumn + 1

---
<img height="200" width="200" src="screenshots/3.png" /><br />
maxColumns * (maxRows - currentRow) + maxColumns - currentColumn + 1

---
<img height="200" width="200" src="screenshots/4.png" /><br />
maxColumns * (maxRows - currentRow) + currentColumn

---
<img height="200" width="200" src="screenshots/5.png" /><br />
maxRows * (currentColumn - 1) + currentRow

---
<img height="200" width="200" src="screenshots/6.png" /><br />
maxRows * (currentColumn - 1) + maxRows - currentRow + 1

---
<img height="200" width="200" src="screenshots/7.png" /><br />
maxRows * (maxColumns - currentColumn) + currentRow

---
<img height="200" width="200" src="screenshots/8.png" /><br />
maxRows * (maxColumns - currentColumn) + maxRows - currentRow + 1

---
<img height="200" width="200" src="screenshots/9.png" /><br />
(int(math.Max(0, math.Pow(-1, float64(currentRow+1)))))*(maxColumns*(currentRow-1)+currentColumn) + ((int(math.Max(0, math.Pow(-1, float64(currentRow))))) * (maxColumns*(currentRow-1) + maxColumns - currentColumn + 1))

---
<img height="200" width="200" src="screenshots/10.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentRow+1))))) * (maxColumns*(currentRow-1) + maxColumns - currentColumn + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentRow)))) * (maxColumns*(currentRow-1) + currentColumn))

---
<img height="200" width="200" src="screenshots/11.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentRow+1))))) * (maxColumns*(maxRows-currentRow) + maxColumns - currentColumn + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentRow)))) * (maxColumns*(maxRows-currentRow) + currentColumn))

---
<img height="200" width="200" src="screenshots/12.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentRow+1))))) * (maxColumns*(maxRows-currentRow) + currentColumn)) + (int(math.Max(0, math.Pow(-1, float64(currentRow)))) * (maxColumns*(maxRows-currentRow) + maxColumns - currentColumn + 1))

---
<img height="200" width="200" src="screenshots/13.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(currentColumn-1) + currentRow)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(currentColumn-1) + maxRows - currentRow + 1))

---
<img height="200" width="200" src="screenshots/14.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(currentColumn-1) + maxRows - currentRow + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(currentColumn-1) + currentRow))

---
<img height="200" width="200" src="screenshots/15.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(maxColumns-currentColumn) + maxRows - currentRow + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(maxColumns-currentColumn) + currentRow))

---
<img height="200" width="200" src="screenshots/16.png" /><br />
((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(maxColumns-currentColumn) + currentRow)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(maxColumns-currentColumn) + maxRows - currentRow + 1))
