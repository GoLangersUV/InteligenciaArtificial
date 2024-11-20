import { Position, Move } from '../types/game';
import { GameUtilities } from './utilities';

export class MoveGenerator {
  static generateMoves(from: Position): Move[] {
    const possiblePositions = GameUtilities.getPossibleMoves(from);
    return possiblePositions.map(to => ({ from, to }));
  }

  static getSortedMovesByProximity(moves: Move[], target: Position): Move[] {
    return [...moves].sort((a, b) => {
      const distA = GameUtilities.getManhattanDistance(a.to, target);
      const distB = GameUtilities.getManhattanDistance(b.to, target);
      return distA - distB;
    });
  }

  static generateSortedMoves(from: Position, target: Position): Move[] {
    const moves = this.generateMoves(from);
    return this.getSortedMovesByProximity(moves, target);
  }
}
