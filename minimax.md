## MiniMax

MiniMax is a backtracking algorithm that alternates between two players: a
maximizer and a minimizer. The maximizer wants the highest score possible while
the minimizer wants the lowest. The algorithm uses an evaluation function to
estimate the value of a given board state. MiniMax assumes that both parties
play optimally.

In order to get the fastest possible win, the score must take the depth into
account (i.e. subtract the depth from the maximizer score and add the depth to
the minimizer score).

One possible evaluation function is to count the number of connected rows,
columns and diagonals for the maximizer.

For example, the following board will:

- O (maximizer) has a score of 4 = 2 (diagonals) + 1 (columns) + 1 (rows)

|   |   |   |   |   |  |   |
|   |   |   | O |   |  |   |
|   |   | O | O |   |  |   |
|   |   | X | X |   |  |   |
|   |   | O | X | X |  |   |
|   | O | X | X | O |  |   |
