const std = @import("std");

const Move = enum(u8) {
    Rock = 0,
    Paper,
    Scissors,
};

const Result = enum(u8) {
    Win = 0,
    Loss,
    Draw,
};

const WIN_POINTS = 6;
const DRAW_POINTS = 3;
const LOSS_POINTS = 0;
const ROCK_POINTS = 1;
const PAPER_POINTS = 2;
const SCISSORS_POINTS = 3;

fn parse_move(char: u8) !Move {
    return switch (char) {
        0, 'A', 'X' => .Rock,
        1, 'B', 'Y' => .Paper,
        2, 'C', 'Z' => .Scissors,
        else => error.UnknownMove,
    };
}

fn parse_result(char: u8) !Result {
    return switch (char) {
        0, 'X' => .Loss,
        1, 'Y' => .Draw,
        2, 'Z' => .Win,
        else => error.UnknownResult,
    };
}

const MOVE_POINTS = [_]u8{ ROCK_POINTS, PAPER_POINTS, SCISSORS_POINTS };

fn game_points(comptime op_move: Move, comptime result: Result) u8 {
    return switch (result) {
        .Win => WIN_POINTS + switch (op_move) {
            .Rock => PAPER_POINTS,
            .Paper => SCISSORS_POINTS,
            .Scissors => ROCK_POINTS,
        },
        .Loss => LOSS_POINTS + switch (op_move) {
            .Rock => SCISSORS_POINTS,
            .Paper => ROCK_POINTS,
            .Scissors => PAPER_POINTS,
        },
        .Draw => DRAW_POINTS + switch (op_move) {
            .Rock => ROCK_POINTS,
            .Paper => PAPER_POINTS,
            .Scissors => SCISSORS_POINTS,
        },
    };
}

fn generate_table() [9]u8 {
    var array: [9]u8 = undefined;
    var i: u8 = 0;
    while (i < 3) : (i += 1) {
        var j: u8 = 0;
        while (j < 3) : (j += 1) {
            const idx = j + 3 * i;
            array[idx] = game_points(parse_move(i) catch .Rock, parse_result(j) catch .Loss);
        }
    }
    return array;
}

const table = generate_table();

pub fn main() !void {
    var file = try std.fs.cwd().openFile("../input/02/question.txt", .{});
    defer file.close();

    var reader = std.io.bufferedReader(file.reader());
    var stream = reader.reader();
    var buffer: [1024]u8 = undefined;

    var score: u32 = 0;
    while (try stream.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        const op_move = @intFromEnum(try parse_move(line[0]));
        const my_move = @intFromEnum(try parse_move(line[2]));

        score += table[my_move + 3 * op_move];
    }

    std.debug.print("{}\n", .{score});
}
