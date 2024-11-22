import { GameState, Position, Horse, BoardValue } from '../types/game';
import { BOARD_SIZE, TOTAL_POINTS_SQUARES, TOTAL_MULTIPLIER_SQUARES } from '../constants/gameConstants';
import { GameUtilities } from './utilities';

export class GameStateManager implements GameState {
  board: BoardValue[][];
  whiteHorse: Horse;
  blackHorse: Horse;
  whiteScore: number;
  blackScore: number;
  currentPlayer: 'white' | 'black';

  constructor(startingPlayer: 'white' | 'black') {
    this.board = this.createEmptyBoard();
    this.whiteHorse = { position: { row: 0, col: 0 }, hasMultiplier: false };
    this.blackHorse = { position: { row: 7, col: 7 }, hasMultiplier: false };
    this.whiteScore = 0;
    this.blackScore = 0;
    this.currentPlayer = startingPlayer;
    this.initializeRandomBoard();
  }

  // Métodos privados de utilidad
  private createEmptyBoard(): BoardValue[][] {
    return Array(BOARD_SIZE).fill(null).map(() => 
      Array(BOARD_SIZE).fill(null).map(() => ({}))
    );
  }

  private clearBoard(): void {
    for (let row = 0; row < BOARD_SIZE; row++) {
      for (let col = 0; col < BOARD_SIZE; col++) {
        this.board[row][col] = {};
      }
    }
  }

  private getCurrentHorse(): Horse {
    return this.currentPlayer === 'white' ? this.whiteHorse : this.blackHorse;
  }

  private processPoints(position: Position, currentHorse: Horse): void {
    const points = this.board[position.row][position.col].points;
    if (points !== undefined) {
      const multiplier = currentHorse.hasMultiplier ? 2 : 1;
      if (this.currentPlayer === 'white') {
        this.whiteScore += points * multiplier;
      } else {
        this.blackScore += points * multiplier;
      }
      this.board[position.row][position.col].points = undefined;
      currentHorse.hasMultiplier = false;
    }
  }

  private processMultiplier(position: Position, currentHorse: Horse): void {
    if (this.board[position.row][position.col].multiplier && !currentHorse.hasMultiplier) {
      currentHorse.hasMultiplier = true;
      this.board[position.row][position.col].multiplier = undefined;
    }
  }

  // Métodos públicos
  initializeRandomBoard(): void {
    this.clearBoard();

    const positions = GameUtilities.getUniqueRandomPositions(
      TOTAL_POINTS_SQUARES + TOTAL_MULTIPLIER_SQUARES + 2
    );

    // Colocar puntos
    for (let i = 0; i < TOTAL_POINTS_SQUARES; i++) {
      const pos = positions[i];
      this.board[pos.row][pos.col].points = i + 1;
    }

    // Colocar multiplicadores
    for (let i = 0; i < TOTAL_MULTIPLIER_SQUARES; i++) {
      const pos = positions[TOTAL_POINTS_SQUARES + i];
      this.board[pos.row][pos.col].multiplier = true;
    }

    // Colocar caballos
    const [whitePos, blackPos] = positions.slice(-2);
    
    this.whiteHorse.position = whitePos;
    this.whiteHorse.hasMultiplier = false;
    this.blackHorse.position = blackPos;
    this.blackHorse.hasMultiplier = false;
    
    this.board[whitePos.row][whitePos.col].horse = 'white';
    this.board[blackPos.row][blackPos.col].horse = 'black';
  }

  isValidMove(from: Position, to: Position): boolean {
    if (!this.board[to.row] || !this.board[to.row][to.col]) {
      return false;
    }

    if (this.board[to.row][to.col].horse) {
      return false;
    }

    return GameUtilities.isValidKnightMove(from, to);
  }

  makeMove(from: Position, to: Position): boolean {
    if (!this.isValidMove(from, to)) {
      return false;
    }

    const currentHorse = this.getCurrentHorse();
    
    if (from.row !== currentHorse.position.row || from.col !== currentHorse.position.col) {
      return false;
    }

    // Realizar el movimiento
    this.board[from.row][from.col].horse = undefined;
    this.board[to.row][to.col].horse = this.currentPlayer;
    currentHorse.position = { ...to };

    // Procesar efectos
    this.processPoints(to, currentHorse);
    this.processMultiplier(to, currentHorse);

    // Cambiar turno
    this.currentPlayer = this.currentPlayer === 'white' ? 'black' : 'white';
    
    return true;
  }

  getPossibleMoves(from: Position): Position[] {
    return GameUtilities.getPossibleMoves(from)
      .filter(pos => !this.board[pos.row][pos.col].horse);
  }

  clone(): GameStateManager {
    const newState = new GameStateManager(this.currentPlayer);
    newState.board.forEach((row, i) => {
      row.forEach((_, j) => {
        newState.board[i][j] = { ...this.board[i][j] };
      });
    });
    newState.whiteHorse.position = { ...this.whiteHorse.position };
    newState.whiteHorse.hasMultiplier = this.whiteHorse.hasMultiplier;
    newState.blackHorse.position = { ...this.blackHorse.position };
    newState.blackHorse.hasMultiplier = this.blackHorse.hasMultiplier;
    newState.whiteScore = this.whiteScore;
    newState.blackScore = this.blackScore;
    newState.currentPlayer = this.currentPlayer;
    return newState;
  }

  hasPointsRemaining(): boolean {
    return this.board.some(row => row.some(cell => cell.points !== undefined));
  }
}
