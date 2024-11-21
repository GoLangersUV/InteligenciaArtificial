import { GameStateManager } from '../gameState';
import { GameUtilities } from '../utilities';
import { Position } from '../../types/game';

export function evaluatePositionAI2(state: GameStateManager): number {
  // Evaluar desde la perspectiva de las blancas
  const baseScore = state.whiteScore - state.blackScore;  // Invertido respecto a AI1
  let positionScore = 0;
  
  const distanceCache = new Map<string, number>();
  const getDistanceFromCache = (from: Position, to: Position): number => {
    const key = `${from.row},${from.col}-${to.row},${to.col}`;
    if (!distanceCache.has(key)) {
      distanceCache.set(key, GameUtilities.getManhattanDistance(from, to));
    }
    return distanceCache.get(key)!;
  };
  
  const whiteMoves = state.getPossibleMoves(state.whiteHorse.position);
  const blackMoves = state.getPossibleMoves(state.blackHorse.position);
  
  // Evaluar movilidad desde perspectiva de blancas
  positionScore += whiteMoves.length * 0.3;  // Premiar nuestra movilidad (blancas)
  positionScore -= blackMoves.length * 0.3;  // Penalizar movilidad del oponente

  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      if (state.board[row][col].points) {
        const points = state.board[row][col].points!;
        const distanceToWhite = getDistanceFromCache(
          state.whiteHorse.position,
          { row, col }
        );
        const distanceToBlack = getDistanceFromCache(
          state.blackHorse.position,
          { row, col }
        );
        
        // Evaluar desde perspectiva de blancas
        if (distanceToWhite < distanceToBlack) {
          // Si blancas están más cerca
          positionScore += points * (4 / (distanceToWhite + 1));
        } else {
          // Si negras están más cerca
          positionScore -= points * (3 / (distanceToBlack + 1));
        }
      }
    }
  }

  // Evaluar amenazas desde perspectiva de blancas
  for (const move of blackMoves) {
    if (state.board[move.row][move.col].points) {
      // Amenaza inmediata de que el oponente (negras) capture puntos
      positionScore -= state.board[move.row][move.col].points! * 2;
    }
  }

  // Evaluar multiplicadores desde perspectiva de blancas
  if (state.whiteHorse.hasMultiplier) {
    positionScore += 15;  // Bonus por tener multiplicador
  }
  if (state.blackHorse.hasMultiplier) {
    positionScore -= 20;  // Penalización si el oponente tiene multiplicador
  }

  let totalPointsRemaining = 0;
  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      if (state.board[row][col].points) {
        totalPointsRemaining += state.board[row][col].points!;
      }
    }
  }
  
  const endgameFactor = totalPointsRemaining > 0 ? 
    totalPointsRemaining / 55 : 
    0.1;

  return (baseScore * 5 * endgameFactor) + positionScore;
}
