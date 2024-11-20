import React from 'react';
import Square from './Square';
import { GameState, Position } from '../../types/game';
import {BOARD_SIZE} from '../../constants/gameConstants';

interface BoardProps {
gameState: GameState;
onClickSquare: (position: Position) => void;
selectedSquare: Position | null;
}

const Board: React.FC<BoardProps> = ({ gameState, onClickSquare, selectedSquare }) => {
  return (
    <div className="grid grid-cols-8 gap-0 border border-gray-400">
      {Array(BOARD_SIZE).fill(null).map((_, row) =>
        Array(BOARD_SIZE).fill(null).map((_, col) => (
          <Square
            key={`${row}-${col}`}
            position={{ row, col }}
            value={gameState.board[row][col]}
            onClick={() => onClickSquare({ row, col })}
            isSelected={selectedSquare?.row === row && selectedSquare?.col === col}
          />
        ))
      )}
    </div>
  );
};

export default Board;
