const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("../input/01/question.txt", .{});
    defer file.close();

    var reader = std.io.bufferedReader(file.reader());
    var stream = reader.reader();
    var buffer: [1024]u8 = undefined;

    var best: u32 = 0;
    var current: u32 = 0;

    while (try stream.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        const end = if (line[line.len-1] == '\r') line.len - 1 else line.len;
        const value = std.fmt.parseInt(u32, line[0..end], 10);

        if (value) |val| {
            current += val;
        } else |_| {
            if (current > best) {
                best = current;
            }

            current = 0;
        }
    }

    std.log.info("{d}", .{best});
} 
