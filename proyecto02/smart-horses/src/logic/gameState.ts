import { GameState, Position, Horse, BoardValue } from '../types/game';
import { BOARD_SIZE, TOTAL_POINTS_SQUARES, TOTAL_MULTIPLIER_SQUARES, POINTS_RANGE, KNIGHT_MOVES } from '../constants/gameConstants';

export class GameStateManager implements GameState {
  board: BoardValue[][];
  whiteHorse: Horse;
  blackHorse: Horse;
  whiteScore: number;
  blackScore: number;
  currentPlayer: 'white' | 'black';

  constructor() {
    this.board = Array(BOARD_SIZE).fill(null).map(() => 
      Array(BOARD_SIZE).fill(null).map(() => ({}))
    );
    this.whiteScore = 0;
    this.blackScore = 0;
    this.currentPlayer = 'white';
    this.whiteHorse = { position: { row: 0, col: 0 }, hasMultiplier: false };
    this.blackHorse = { position: { row: 7, col: 7 }, hasMultiplier: false };
    this.initializeRandomBoard();
  }

  initializeRandomBoard(): void {
    // Limpiar el tablero
    this.board = Array(BOARD_SIZE).fill(null).map(() => 
      Array(BOARD_SIZE).fill(null).map(() => ({}))
    );

    const availablePositions: Position[] = [];
    for (let row = 0; row < BOARD_SIZE; row++) {
      for (let col = 0; col < BOARD_SIZE; col++) {
        availablePositions.push({ row, col });
      }
    }

    // Función para obtener una posición aleatoria
    const getRandomPosition = (): Position => {
      const index = Math.floor(Math.random() * availablePositions.length);
      return availablePositions.splice(index, 1)[0];
    };

    // Colocar puntos
    for (let i = 1; i <= TOTAL_POINTS_SQUARES; i++) {
      const pos = getRandomPosition();
      this.board[pos.row][pos.col].points = i;
    }

    // Colocar multiplicadores
    for (let i = 0; i < TOTAL_MULTIPLIER_SQUARES; i++) {
      const pos = getRandomPosition();
      this.board[pos.row][pos.col].multiplier = true;
    }

    // Colocar caballos
    const whitePos = getRandomPosition();
    const blackPos = getRandomPosition();
    
    this.whiteHorse = { position: whitePos, hasMultiplier: false };
    this.blackHorse = { position: blackPos, hasMultiplier: false };
    
    this.board[whitePos.row][whitePos.col].horse = 'white';
    this.board[blackPos.row][blackPos.col].horse = 'black';
  }

  isValidMove(from: Position, to: Position): boolean {
    // Verificar si está dentro del tablero
    if (to.row < 0 || to.row >= BOARD_SIZE || to.col < 0 || to.col >= BOARD_SIZE) {
      return false;
    }

    // Verificar movimiento en L
    const rowDiff = Math.abs(to.row - from.row);
    const colDiff = Math.abs(to.col - from.col);
    return (rowDiff === 2 && colDiff === 1) || (rowDiff === 1 && colDiff === 2);
  }

  makeMove(from: Position, to: Position): boolean {
    if (!this.isValidMove(from, to)) return false;

    const horse = this.currentPlayer === 'white' ? this.whiteHorse : this.blackHorse;
    
    // Limpiar posición anterior
    this.board[from.row][from.col].horse = undefined;

    // Mover caballo
    horse.position = to;
    this.board[to.row][to.col].horse = this.currentPlayer;

    // Procesar puntos y multiplicadores
    if (this.board[to.row][to.col].points) {
      const points = this.board[to.row][to.col].points!;
      const multiplier = horse.hasMultiplier ? 2 : 1;
      
      if (this.currentPlayer === 'white') {
        this.whiteScore += points * multiplier;
      } else {
        this.blackScore += points * multiplier;
      }

      this.board[to.row][to.col].points = undefined;
      horse.hasMultiplier = false;
    }

    if (this.board[to.row][to.col].multiplier) {
      horse.hasMultiplier = true;
      this.board[to.row][to.col].multiplier = undefined;
    }

    this.currentPlayer = this.currentPlayer === 'white' ? 'black' : 'white';
    return true;
  }

  clone(): GameStateManager {
    const newState = new GameStateManager();
    newState.board = JSON.parse(JSON.stringify(this.board));
    newState.whiteHorse = { ...this.whiteHorse };
    newState.blackHorse = { ...this.blackHorse };
    newState.whiteScore = this.whiteScore;
    newState.blackScore = this.blackScore;
    newState.currentPlayer = this.currentPlayer;
    return newState;
  }

  getPossibleMoves(from: Position): Position[] {
    const moves: Position[] = [];
    
    for (const [rowDiff, colDiff] of KNIGHT_MOVES) {
      const newRow = from.row + rowDiff;
      const newCol = from.col + colDiff;
      
      if (newRow >= 0 && newRow < BOARD_SIZE && newCol >= 0 && newCol < BOARD_SIZE) {
        moves.push({ row: newRow, col: newCol });
      }
    }
    
    return moves;
  }
}
