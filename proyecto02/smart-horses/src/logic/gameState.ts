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
    this.board = Array(BOARD_SIZE).fill(null).map(() => 
      Array(BOARD_SIZE).fill(null).map(() => ({}))
    );

    const positions = GameUtilities.getUniqueRandomPositions(
      TOTAL_POINTS_SQUARES + TOTAL_MULTIPLIER_SQUARES + 2 // +2 para los caballos
    );

    for (let i = 0; i < TOTAL_POINTS_SQUARES; i++) {
      const pos = positions[i];
      this.board[pos.row][pos.col].points = i + 1;
    }

    for (let i = 0; i < TOTAL_MULTIPLIER_SQUARES; i++) {
      const pos = positions[TOTAL_POINTS_SQUARES + i];
      this.board[pos.row][pos.col].multiplier = true;
    }

    const whitePos = positions[positions.length - 2];
    const blackPos = positions[positions.length - 1];
    
    this.whiteHorse = { position: whitePos, hasMultiplier: false };
    this.blackHorse = { position: blackPos, hasMultiplier: false };
    
    this.board[whitePos.row][whitePos.col].horse = 'white';
    this.board[blackPos.row][blackPos.col].horse = 'black';
  }

  isValidMove(from: Position, to: Position): boolean {
    if (this.board[to.row]?.[to.col]?.horse) {
      return false;
    }

    return GameUtilities.isValidKnightMove(from, to);
  }

  makeMove(from: Position, to: Position): boolean {
    if (!this.isValidMove(from, to)) {
      return false;
    }

    const currentHorse = this.currentPlayer === 'white' ? this.whiteHorse : this.blackHorse;
    
    if (from.row !== currentHorse.position.row || from.col !== currentHorse.position.col) {
      return false;
    }

    // Realizar el movimiento
    this.board[from.row][from.col].horse = undefined;
    this.board[to.row][to.col].horse = this.currentPlayer;
    currentHorse.position = { ...to };

    // Procesar puntos si los hay
    if (this.board[to.row][to.col].points) {
      const points = this.board[to.row][to.col].points!;
      const multiplier = currentHorse.hasMultiplier ? 2 : 1;
      
      if (this.currentPlayer === 'white') {
        this.whiteScore += points * multiplier;
      } else {
        this.blackScore += points * multiplier;
      }

      this.board[to.row][to.col].points = undefined;
      currentHorse.hasMultiplier = false;
    }

    // Procesar multiplicador
    if (this.board[to.row][to.col].multiplier && !currentHorse.hasMultiplier) {
      currentHorse.hasMultiplier = true;
      this.board[to.row][to.col].multiplier = undefined;
    }

    this.currentPlayer = this.currentPlayer === 'white' ? 'black' : 'white';
    return true;
  }

  getPossibleMoves(from: Position): Position[] {
    return GameUtilities.getPossibleMoves(from)
      .filter(pos => !this.board[pos.row][pos.col].horse);
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

  hasPointsRemaining(): boolean {
    return this.board.some(row => row.some(cell => cell.points !== undefined));
  }
}
