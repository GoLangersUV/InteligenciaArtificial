import { GameStateManager } from '../gameState';
import { GameUtilities } from '../utilities';
import { Position } from '../../types/game';

export function evaluatePositionAI1(state: GameStateManager): number {
  const baseScore = state.whiteScore - state.blackScore;
  let positionScore = 0;
  
  // Cache de distancias para evitar recálculos
  const distanceCache = new Map<string, number>();
  
  const getDistanceFromCache = (from: Position, to: Position): number => {
    const key = `${from.row},${from.col}-${to.row},${to.col}`;
    if (!distanceCache.has(key)) {
      distanceCache.set(key, GameUtilities.getManhattanDistance(from, to));
    }
    return distanceCache.get(key)!;
  };

  // Evaluación de puntos y posición
  let totalPointsRemaining = 0;
  let bestPointValue = 0;
  
  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      if (state.board[row][col].points) {
        const points = state.board[row][col].points!;
        totalPointsRemaining += points;
        bestPointValue = Math.max(bestPointValue, points);
        
        const distanceToWhite = getDistanceFromCache(
          state.whiteHorse.position,
          { row, col }
        );
        const distanceToBlack = getDistanceFromCache(
          state.blackHorse.position,
          { row, col }
        );
        
        // Factor dinámico basado en puntos restantes
        const importanceFactor = points / totalPointsRemaining;
        
        // Evaluación ponderada de la posición
        if (distanceToWhite < distanceToBlack) {
          positionScore += points * (1.5 + importanceFactor);
        } else if (distanceToWhite > distanceToBlack) {
          positionScore -= points * (1 + importanceFactor);
        }
      }
    }
  }

  // Evaluación de multiplicadores
  if (!state.whiteHorse.hasMultiplier) {
    let nearestMultiplierDistance = Infinity;
    let potentialValue = 0;
    
    for (let row = 0; row < 8; row++) {
      for (let col = 0; col < 8; col++) {
        if (state.board[row][col].multiplier) {
          const distance = getDistanceFromCache(
            state.whiteHorse.position,
            { row, col }
          );
          if (distance < nearestMultiplierDistance) {
            nearestMultiplierDistance = distance;
            potentialValue = bestPointValue;
          }
        }
      }
    }
    
    if (nearestMultiplierDistance !== Infinity) {
      positionScore += (potentialValue / nearestMultiplierDistance) * 1.5;
    }
  }

  // Evaluación de movilidad y control
  const whiteMoves = state.getPossibleMoves(state.whiteHorse.position);
  const mobilityScore = whiteMoves.length * 0.2;
  
  // Factor de fin de juego
  const endgameFactor = totalPointsRemaining > 0 ? 
    totalPointsRemaining / 55 :
    0.1;
  
  return (baseScore * 3 * endgameFactor) + (positionScore * 1.5) + mobilityScore;
}
