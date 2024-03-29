const std = @import("std");

const Move = enum(u8) {
    Rock = 0,
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
        0, 'A', 'X' => .Rock,
        1, 'B', 'Y' => .Paper,
        2, 'C', 'Z' => .Scissors,
        else => error.UnknownMove,
    };
}

const MOVE_POINTS = [_]u8{ ROCK_POINTS, PAPER_POINTS, SCISSORS_POINTS };

fn game_points(comptime i: Move, comptime j: Move) u8 {
    return switch (i) {
        .Rock => switch (j) {
            .Rock => DRAW_POINTS,
            .Paper => WIN_POINTS,
            .Scissors => LOSS_POINTS,
        },
        .Paper => switch (j) {
            .Rock => LOSS_POINTS,
            .Paper => DRAW_POINTS,
            .Scissors => WIN_POINTS,
        },
        .Scissors => switch (j) {
            .Rock => WIN_POINTS,
            .Paper => LOSS_POINTS,
            .Scissors => DRAW_POINTS,
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
            array[idx] = MOVE_POINTS[j] + game_points(parse_move(i) catch .Rock, parse_move(j) catch .Rock);
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
