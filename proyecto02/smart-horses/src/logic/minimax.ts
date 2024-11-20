// src/logic/minimax.ts

import { GameStateManager } from './gameState';
import { Move } from '../types/game';

export class Minimax {
  private evaluationFn: (state: GameStateManager) => number;
  private maxDepth: number;

  constructor(evaluationFn: (state: GameStateManager) => number, maxDepth: number) {
    this.evaluationFn = evaluationFn;
    this.maxDepth = maxDepth;
  }

  getBestMove(state: GameStateManager): Move {
        console.log('Getting best move for state:', {
            currentPlayer: state.currentPlayer,
            whiteHorse: state.whiteHorse,
            blackHorse: state.blackHorse
        });

        const currentHorse = state.currentPlayer === 'white' ? state.whiteHorse : state.blackHorse;
        const possibleMoves = state.getPossibleMoves(currentHorse.position);
        
        console.log('Possible moves:', possibleMoves);

        let bestScore = -Infinity;
        let bestMove: Move = {
            from: currentHorse.position,
            to: possibleMoves[0]
        };

        for (const to of possibleMoves) {
            const newState = state.clone();
            const success = newState.makeMove(currentHorse.position, to);
            
            if (success) {
                const score = this.minimax(newState, this.maxDepth - 1, false, -Infinity, Infinity);
                console.log('Move evaluation:', { to, score });
                
                if (score > bestScore) {
                    bestScore = score;
                    bestMove = {
                        from: currentHorse.position,
                        to: to
                    };
                }
            }
        }

        console.log('Selected best move:', bestMove);
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

    const currentHorse = isMaximizing ? state.whiteHorse : state.blackHorse;
    const possibleMoves = state.getPossibleMoves(currentHorse.position);

    if (isMaximizing) {
      let maxEval = -Infinity;
      
      for (const to of possibleMoves) {
        const newState = state.clone();
        if (newState.makeMove(currentHorse.position, to)) {
          const eval_ = this.minimax(newState, depth - 1, false, alpha, beta);
          maxEval = Math.max(maxEval, eval_);
          alpha = Math.max(alpha, eval_);
          
          if (beta <= alpha) {
            break;
          }
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
          
          if (beta <= alpha) {
            break;
          }
        }
      }
      
      return minEval;
    }
  }
}
