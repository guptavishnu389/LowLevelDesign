# Snake and Ladder Game - Problem Statement

## Game Overview
Snake and Ladder is a classic board game where players move their pieces according to dice rolls, climbing ladders and sliding down snakes to reach the final position (usually 100).

## Requirements

### 1. Board Setup
- The game board consists of 100 squares (1 to 100)
- Snakes and ladders are placed at specific positions
- Each snake has a head (start) and tail (end)
- Each ladder has a bottom (start) and top (end)

### 2. Game Components
- Multiple players (2 or more)
- A six-sided die
- Game board with snakes and ladders
- Player pieces/tokens

### 3. Game Rules
1. Players take turns to roll the die and move their piece
2. Movement is based on the number rolled on the die
3. If a player lands on:
   - A ladder's bottom: Move up to the ladder's top
   - A snake's head: Slide down to the snake's tail
4. To win, a player must reach exactly square 100
   - If a roll would take the player beyond 100, they stay in their current position
   - The player must roll the exact number needed to reach 100

### 4. Technical Requirements
1. Support for multiple players
2. Random dice roll generation
3. Snake and ladder position management
4. Player position tracking
5. Game state management
6. Win condition checking

### 5. Implementation Details
- Use object-oriented design principles
- Implement proper error handling
- Follow clean code practices
- Make the code extensible for future modifications

### 6. Sample Snake and Ladder Positions
```
Snakes:
- 16 -> 6
- 47 -> 26
- 49 -> 11
- 56 -> 53
- 62 -> 19
- 64 -> 60
- 87 -> 24
- 93 -> 73
- 95 -> 75
- 98 -> 78

Ladders:
- 1 -> 38
- 4 -> 14
- 9 -> 31
- 21 -> 42
- 28 -> 84
- 36 -> 44
- 51 -> 67
- 71 -> 91
- 80 -> 100
```

## Expected Output
The game should:
1. Initialize with the specified number of players
2. Display the current state of the game after each move
3. Show which player's turn it is
4. Display the dice roll result
5. Show the player's movement and any snake/ladder encounters
6. Announce the winner when a player reaches square 100 