const std = @import("std");
const math = std.math;

fn compare(context: void, a: u32, b: u32) math.Order {
    _ = context;
    return math.order(b, a);
}

const PQueue = std.PriorityQueue(u32, void, compare);

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    var queue = PQueue.init(allocator, {});
    defer queue.deinit();

    var file = try std.fs.cwd().openFile("../input/01/question.txt", .{});
    defer file.close();

    var reader = std.io.bufferedReader(file.reader());
    var stream = reader.reader();
    var buffer: [1024]u8 = undefined;

    var current: u32 = 0;

    while (try stream.readUntilDelimiterOrEof(&buffer, '\n')) |line| {
        const end = if (line[line.len - 1] == '\r') line.len - 1 else line.len;
        const value = std.fmt.parseInt(u32, line[0..end], 10);

        if (value) |val| {
            current += val;
        } else |_| {
            try queue.add(current);
            current = 0;
        }
    }
    try queue.add(current);

    const answer = queue.remove() + queue.remove() + queue.remove();
    std.debug.print("{}", .{answer});
}
