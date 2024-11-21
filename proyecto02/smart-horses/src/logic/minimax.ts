import { GameStateManager } from './gameState';
import { Move, Position } from '../types/game';

export class Minimax {
  private evaluationFn: (state: GameStateManager) => number;
  private maxDepth: number;
  private moveCounter: Map<string, number>;
  private readonly MAX_REPETITIONS = 2;

  constructor(evaluationFn: (state: GameStateManager) => number, maxDepth: number) {
    this.evaluationFn = evaluationFn;
    this.maxDepth = maxDepth;
    this.moveCounter = new Map<string, number>();
  }

  private getMoveKey(from: Position, to: Position): string {
    return `${from.row},${from.col}->${to.row},${to.col}`;
  }

  private incrementMoveCount(from: Position, to: Position): number {
    const moveKey = this.getMoveKey(from, to);
    const currentCount = this.moveCounter.get(moveKey) || 0;
    this.moveCounter.set(moveKey, currentCount + 1);
    return currentCount + 1;
  }

  private getRepetitionCount(from: Position, to: Position): number {
    const moveKey = this.getMoveKey(from, to);
    return this.moveCounter.get(moveKey) || 0;
  }

  getBestMove(state: GameStateManager): Move | null {
    console.log('Getting best move for state:', {
      currentPlayer: state.currentPlayer,
      whiteHorse: state.whiteHorse,
      blackHorse: state.blackHorse
    });

    if (state.currentPlayer !== 'black') {
      console.log('Not AI turn');
      return null;
    }

    const currentHorse = state.blackHorse;
    let possibleMoves = state.getPossibleMoves(currentHorse.position);
    
    console.log('Possible moves:', possibleMoves);

    if (possibleMoves.length === 0) {
      console.log('No possible moves');
      return null;
    }

    let validMoves = possibleMoves.filter(to => {
      const repetitions = this.getRepetitionCount(currentHorse.position, to);
      return repetitions < this.MAX_REPETITIONS;
    });

    if (validMoves.length === 0) {
      console.log('Resetting move counters due to all moves being repeated');
      this.moveCounter.clear();
      validMoves = possibleMoves;
    }

    let bestScore = -Infinity;
    let bestMove: Move | null = null;

    for (const to of validMoves) {
      const newState = state.clone();
      const success = newState.makeMove(currentHorse.position, to);
      
      if (success) {
        const score = this.minimax(newState, this.maxDepth - 1, false, -Infinity, Infinity);
        
        const immediatePoints = state.board[to.row][to.col].points || 0;
        const hasMultiplier = state.board[to.row][to.col].multiplier;
        const immediateBonus = (immediatePoints * 1000) + (hasMultiplier ? 500 : 0);
        
        const finalScore = score + immediateBonus;
        
        console.log('Move evaluation:', { 
          to, 
          score: finalScore,
          repetitions: this.getRepetitionCount(currentHorse.position, to)
        });

        if (finalScore > bestScore) {
          bestScore = finalScore;
          bestMove = {
            from: currentHorse.position,
            to: to
          };
        }
      }
    }

    if (bestMove) {
      this.incrementMoveCount(bestMove.from, bestMove.to);
      console.log('Selected move:', bestMove, 
        'Times used:', this.getRepetitionCount(bestMove.from, bestMove.to));
    } else if (possibleMoves.length > 0) {
      bestMove = {
        from: currentHorse.position,
        to: possibleMoves[0]
      };
      console.log('Fallback to first available move:', bestMove);
    }

    return bestMove;
  }

  private minimax(
    state: GameStateManager,
    depth: number,
    isMaximizing: boolean,
    alpha: number,
    beta: number
  ): number {
    if (depth === 0 || !state.hasPointsRemaining()) {
      return this.evaluationFn(state);
    }

    const currentHorse = isMaximizing ? state.blackHorse : state.whiteHorse;
    const possibleMoves = state.getPossibleMoves(currentHorse.position);

    if (possibleMoves.length === 0) {
      return isMaximizing ? -Infinity : Infinity;
    }

    if (isMaximizing) { 
      let maxEval = -Infinity;
      
      for (const to of possibleMoves) {
        const newState = state.clone();
        if (newState.makeMove(currentHorse.position, to)) {
          const eval_ = this.minimax(newState, depth - 1, false, alpha, beta);
          maxEval = Math.max(maxEval, eval_);
          alpha = Math.max(alpha, eval_);
          
          if (beta <= alpha) break;
        }
      }
      
      return maxEval;
    } else {
      let minEval = Infinity;
      
      for (const to of possibleMoves) {
        const newState = state.clone();
        if (newState.makeMove(currentHorse.position, to)) {
          const eval_ = this.minimax(newState, depth - 1, true, alpha, beta);
          minEval = Math.min(minEval, eval_);
          beta = Math.min(beta, eval_);
          
          if (beta <= alpha) break;
        }
      }
      
      return minEval;
    }
  }
}
