const std = @import("std");

const Move = enum {
    Rock,
    Paper,
    Scissors,
};

const WIN_POINTS = 6;
const DRAW_POINTS = 3;
const LOSS_POINTS = 0;
const ROCK_POINTS = 1;
const PAPER_POINTS = 2;
const SCISSORS_POINTS = 3;

fn parse_move(char: u8) !Move {
    return switch (char) {
        'A', 'X' => .Rock,
        'B', 'Y' => .Paper,
        'C', 'Z' => .Scissors,
        else => error.UnknownMoveChar,
    };
}

pub fn main() !void {
    var file = try std.fs.cwd().openFile("../input/02/question.txt", .{});
    defer file.close();

    var reader = std.io.bufferedReader(file.reader());
    var stream = reader.reader();
    var buffer: [1024]u8 = undefined;

    var score: u32 = 0;
    while (try stream.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        const op_move = try parse_move(line[0]);
        const my_move = try parse_move(line[2]);

        switch (my_move) {
            .Rock => {
                score += ROCK_POINTS;
                score += switch (op_move) {
                    .Rock => DRAW_POINTS,
                    .Paper => LOSS_POINTS,
                    .Scissors => WIN_POINTS,
                };
            },
            .Paper => {
                score += PAPER_POINTS;
                score += switch (op_move) {
                    .Rock => WIN_POINTS,
                    .Paper => DRAW_POINTS,
                    .Scissors => LOSS_POINTS,
                };
            },
            .Scissors => {
                score += SCISSORS_POINTS;
                score += switch (op_move) {
                    .Rock => LOSS_POINTS,
                    .Paper => WIN_POINTS,
                    .Scissors => DRAW_POINTS,
                };
            },
        }
    }
}
