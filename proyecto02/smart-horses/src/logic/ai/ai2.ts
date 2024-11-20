import { GameStateManager } from '../gameState';
import { GameUtilities } from '../utilities';

export function evaluatePositionAI2(state: GameStateManager): number {
  const baseScore = state.whiteScore - state.blackScore;
  let positionScore = 0;
  
  const whiteMoves = state.getPossibleMoves(state.whiteHorse.position);
  const blackMoves = state.getPossibleMoves(state.blackHorse.position);
  
  positionScore += whiteMoves.length * 0.5;
  positionScore -= blackMoves.length * 0.7;

  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      if (state.board[row][col].points) {
        const points = state.board[row][col].points!;
        const distanceToWhite = GameUtilities.getManhattanDistance(
          state.whiteHorse.position,
          { row, col }
        );
        const distanceToBlack = GameUtilities.getManhattanDistance(
          state.blackHorse.position,
          { row, col }
        );
        
        if (distanceToWhite < distanceToBlack) {
          positionScore += points;
        } else {
          positionScore -= points * 0.5;
        }
      }
    }
  }

  for (const move of blackMoves) {
    if (state.board[move.row][move.col].points) {
      positionScore -= state.board[move.row][move.col].points! * 0.8;
    }
  }

  if (state.whiteHorse.hasMultiplier) {
    positionScore += 5;
  }
  if (state.blackHorse.hasMultiplier) {
    positionScore -= 8;
  }

  return baseScore * 2 + positionScore;
}
