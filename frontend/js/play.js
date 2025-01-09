document.addEventListener("DOMContentLoaded", () => {
    const rows = 12;
    const cols = 12;
    const minesCount = 20;
    let mineField = [];
    let revealedCount = 0;
    let gameOver = false;
    let isCtrlPressed = false;
    let gameStarted = false; 
    let moveCount = 0; 

    const gameBoard = document.getElementById("game-board");
    const restartButton = document.getElementById("restart-button");

    restartButton.addEventListener("click", () => {
        location.reload();
    });

    function initField() {
        mineField = Array(rows).fill().map(() => Array(cols).fill(null).map(() => ({ mine: false, revealed: false, flagged: false, adjacentMines: 0 })));

        let minesPlaced = 0;
        while (minesPlaced < minesCount) {
            const row = Math.floor(Math.random() * rows);
            const col = Math.floor(Math.random() * cols);
            if (!mineField[row][col].mine) {
                mineField[row][col].mine = true;
                minesPlaced++;
            }
        }

        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (!mineField[row][col].mine) {
                    mineField[row][col].adjacentMines = countAdjacentMines(row, col);
                }
            }
        }
    }

    function countAdjacentMines(row, col) {
        const directions = [
            [-1, -1], [-1, 0], [-1, 1],
            [0, -1],         [0, 1],
            [1, -1], [1, 0], [1, 1]
        ];
        let count = 0;
        directions.forEach(([dx, dy]) => {
            const newRow = row + dx;
            const newCol = col + dy;
            if (newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && mineField[newRow][newCol].mine) {
                count++;
            }
        });
        return count;
    }

    function createBoard() {
        gameBoard.innerHTML = '';
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                const cell = document.createElement("div");
                cell.classList.add("cell");
                cell.dataset.row = row;
                cell.dataset.col = col;
                cell.addEventListener("click", handleLeftClick);
                cell.addEventListener("contextmenu", handleRightClick);
                gameBoard.appendChild(cell);
            }
        }
    }

    function handleLeftClick(event) {
        if (gameOver) return;
        const cell = event.target;
        const row = parseInt(cell.dataset.row);
        const col = parseInt(cell.dataset.col);

        if (!gameStarted) {
            gameStarted = true;
            document.dispatchEvent(new CustomEvent('mine.start')); 
        }

        moveCount++; 
        document.dispatchEvent(new CustomEvent('mine.step', { detail: { row, col, moveCount } })); 

        if (mineField[row][col].flagged || mineField[row][col].revealed) return;
        if (mineField[row][col].mine) {
            revealMines();
            endGame(false);
        } else {
            revealCell(row, col);
            if (revealedCount === rows * cols - minesCount) {
                endGame(true);
            }
        }
    }

    function handleRightClick(event) {
        event.preventDefault();
        if (gameOver) return;
        const cell = event.target;
        const row = parseInt(cell.dataset.row);
        const col = parseInt(cell.dataset.col);
        if (mineField[row][col].revealed) return;
        mineField[row][col].flagged = !mineField[row][col].flagged;
        cell.classList.toggle("flagged");
    }

    function revealCell(row, col) {
        const cell = document.querySelector(`[data-row='${row}'][data-col='${col}']`);
        if (mineField[row][col].revealed || mineField[row][col].flagged) return;
        mineField[row][col].revealed = true;
        cell.classList.add("revealed");
        revealedCount++;
        if (mineField[row][col].adjacentMines > 0) {
            cell.textContent = mineField[row][col].adjacentMines;
        } else {
            const directions = [
                [-1, -1], [-1, 0], [-1, 1],
                [0, -1],         [0, 1],
                [1, -1], [1, 0], [1, 1]
            ];
            directions.forEach(([dx, dy]) => {
                const newRow = row + dx;
                const newCol = col + dy;
                if (newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols) {
                    revealCell(newRow, newCol);
                }
            });
        }
    }

    function revealMines() {
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (mineField[row][col].mine) {
                    const cell = document.querySelector(`[data-row='${row}'][data-col='${col}']`);
                    cell.classList.add("revealed");
                    cell.textContent = "💣";
                }
            }
        }
    }

    function endGame(won) {
        gameOver = true;
        const cells = document.querySelectorAll(".cell");
        cells.forEach(cell => cell.classList.add("disabled"));
        restartButton.style.display = "block";

        document.dispatchEvent(new CustomEvent('mine.end', { detail: { won, moveCount } })); 
    }

    let selectedCell = { row: 0, col: 0 };

    function handleKeyDown(event) {
        const { row, col } = selectedCell;
        if (gameOver) return;

        if (event.key === 'Control') {
            isCtrlPressed = true;
            return;
        }

        switch (event.key) {
            case "ArrowUp":
                if (row > 0) {
                    updateSelectedCell(row - 1, col);
                }
                break;
            case "ArrowDown":
                if (row < rows - 1) {
                    updateSelectedCell(row + 1, col);
                }
                break;
            case "ArrowLeft":
                if (col > 0) {
                    updateSelectedCell(row, col - 1);
                }
                break;
            case "ArrowRight":
                if (col < cols - 1) {
                    updateSelectedCell(row, col + 1);
                }
                break;
            case "Enter":
            case " ":
                if (isCtrlPressed) {
                    handleRightClick({
                        target: document.querySelector(`[data-row='${row}'][data-col='${col}']`),
                        preventDefault: () => {}
                    });
                } else {
                    handleLeftClick({
                        target: document.querySelector(`[data-row='${row}'][data-col='${col}']`)
                    });
                }
                break;
        }
    }

    function handleKeyUp(event) {
        if (event.key === 'Control') {
            isCtrlPressed = false;
        }
    }

    function updateSelectedCell(newRow, newCol) {
        const prevCell = document.querySelector(`[data-row='${selectedCell.row}'][data-col='${selectedCell.col}']`);
        prevCell.classList.remove("selected");
        selectedCell = { row: newRow, col: newCol };
        const newCell = document.querySelector(`[data-row='${newRow}'][data-col='${newCol}']`);
        newCell.classList.add("selected");
    }

    document.addEventListener("keydown", handleKeyDown);
    document.addEventListener("keyup", handleKeyUp);
    initField();
    createBoard();

    document.querySelector('[data-row="0"][data-col="0"]').classList.add("selected");

    document.addEventListener('mine.start', () => {
        console.log("Игра началась");
        restartButton.style.display = "none"; 
    });

    document.addEventListener('mine.step', (event) => {
        console.log(`Строка: ${event.detail.row}, Колонка: ${event.detail.col}`);
        console.log(`Шаг: ${event.detail.moveCount}`); 
    });

    document.addEventListener('mine.end', (event) => {
        console.log(`Игра окончена. Победа: ${event.detail.won}`);
        alert(`Игра  окончена. ВЫ ${event.detail.won ? "выиграли" : "проиграли"} на ${event.detail.moveCount} ходу/е.`);
        restartButton.style.display = "block"; 
    });
});
