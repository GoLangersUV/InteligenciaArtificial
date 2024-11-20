import React from 'react';
import Square from './Square';
import { GameState, Position } from '../../types/game';
import { BOARD_SIZE } from '../../constants/gameConstants';

interface BoardProps {
  gameState: GameState;
  onSquareClick: (position: Position) => void;
  selectedSquare: Position | null;
}

const Board: React.FC<BoardProps> = ({ gameState, onSquareClick, selectedSquare }) => {
  return (
    <div className="grid grid-cols-8 w-[640px] h-[640px] mx-auto border-2 border-gray-600">
      {Array(BOARD_SIZE).fill(null).map((_, row) =>
        Array(BOARD_SIZE).fill(null).map((_, col) => (
          <Square
            key={`${row}-${col}`}
            position={{ row, col }}
            value={gameState.board[row][col]}
            onClick={() => onSquareClick({ row, col })}
            isSelected={selectedSquare?.row === row && selectedSquare?.col === col}
          />
        ))
      )}
    </div>
  );
};

export default Board;
