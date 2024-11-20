import { Position, BoardValue, GameState } from '../types/game';
import { BOARD_SIZE, KNIGHT_MOVES } from '../constants/gameConstants';

export class GameUtilities {
  static getRandomPosition(): Position {
    return {
      row: Math.floor(Math.random() * BOARD_SIZE),
      col: Math.floor(Math.random() * BOARD_SIZE)
    };
  }

  static getUniqueRandomPositions(count: number): Position[] {
    const positions: Position[] = [];
    const used = new Set<string>();

    while (positions.length < count) {
      const pos = this.getRandomPosition();
      const key = `${pos.row},${pos.col}`;
      
      if (!used.has(key)) {
        used.add(key);
        positions.push(pos);
      }
    }

    return positions;
  }

  static isValidPosition(position: Position): boolean {
    return (
      position.row >= 0 &&
      position.row < BOARD_SIZE &&
      position.col >= 0 &&
      position.col < BOARD_SIZE
    );
  }

  static isValidKnightMove(from: Position, to: Position): boolean {
    if (!this.isValidPosition(to)) return false;
    
    const rowDiff = Math.abs(to.row - from.row);
    const colDiff = Math.abs(to.col - from.col);
    return (rowDiff === 2 && colDiff === 1) || (rowDiff === 1 && colDiff === 2);
  }

  static getPossibleMoves(from: Position): Position[] {
    return KNIGHT_MOVES
      .map(([rowDiff, colDiff]) => ({
        row: from.row + rowDiff,
        col: from.col + colDiff
      }))
      .filter(pos => this.isValidPosition(pos));
  }

  static getManhattanDistance(pos1: Position, pos2: Position): number {
    return Math.abs(pos1.row - pos2.row) + Math.abs(pos1.col - pos2.col);
  }

  static getChebyshevDistance(pos1: Position, pos2: Position): number {
    return Math.max(Math.abs(pos1.row - pos2.row), Math.abs(pos1.col - pos2.col));
  }

  static findHighestPointOnBoard(board: BoardValue[][]): { position: Position; points: number } | null {
    let highest = { position: { row: 0, col: 0 }, points: 0 };
    
    for (let row = 0; row < BOARD_SIZE; row++) {
      for (let col = 0; col < BOARD_SIZE; col++) {
        if (board[row][col].points && board[row][col].points! > highest.points) {
          highest = {
            position: { row, col },
            points: board[row][col].points!
          };
        }
      }
    }

    return highest.points > 0 ? highest : null;
  }

  static calculateTotalRemainingPoints(board: BoardValue[][]): number {
    let total = 0;
    for (let row = 0; row < BOARD_SIZE; row++) {
      for (let col = 0; col < BOARD_SIZE; col++) {
        if (board[row][col].points) {
          total += board[row][col].points!;
        }
      }
    }
    return total;
  }

  static evaluatePosition(position: Position, gameState: GameState): number {
    let score = 0;
    const pointsWeight = 1.5;
    const multiplierWeight = 1.0;
    const centerWeight = 0.5;

    // Evaluate proximity to points
    for (let row = 0; row < BOARD_SIZE; row++) {
      for (let col = 0; col < BOARD_SIZE; col++) {
        if (gameState.board[row][col].points) {
          const distance = this.getManhattanDistance(position, { row, col });
          score += (gameState.board[row][col].points! / distance) * pointsWeight;
        }
      }
    }

    // Evaluate proximity to multipliers
    for (let row = 0; row < BOARD_SIZE; row++) {
      for (let col = 0; col < BOARD_SIZE; col++) {
        if (gameState.board[row][col].multiplier) {
          const distance = this.getManhattanDistance(position, { row, col });
          score += (5 / distance) * multiplierWeight;
        }
      }
    }

    // Evaluate center control
    const centerDistance = this.getChebyshevDistance(
      position,
      { row: BOARD_SIZE / 2 - 0.5, col: BOARD_SIZE / 2 - 0.5 }
    );
    score += (1 / (centerDistance + 1)) * centerWeight;

    return score;
  }

  static debugPrintBoard(gameState: GameState): void {
    console.log('\nCurrent Board State:');
    for (let row = 0; row < BOARD_SIZE; row++) {
      let rowStr = '';
      for (let col = 0; col < BOARD_SIZE; col++) {
        const cell = gameState.board[row][col];
        if (cell.horse === 'white') rowStr += '♘ ';
        else if (cell.horse === 'black') rowStr += '♞ ';
        else if (cell.points) rowStr += `${cell.points} `;
        else if (cell.multiplier) rowStr += 'x2 ';
        else rowStr += '· ';
      }
      console.log(rowStr);
    }
    console.log(`White Score: ${gameState.whiteScore}`);
    console.log(`Black Score: ${gameState.blackScore}`);
    console.log(`Current Turn: ${gameState.currentPlayer}`);
  }
}
