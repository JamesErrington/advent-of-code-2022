const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    var arena = std.heap.ArenaAllocator.init(gpa.allocator());
    defer arena.deinit();
    const allocator = arena.allocator();

    const args = try std.process.argsAlloc(allocator);
    if (args.len < 2) {
        std.log.err("No day argument given. Usage: zig run <file.zig> -- [day]", .{});
        std.process.abort();
    }
    const day = args[1];

    const cwd = try std.process.getCwdAlloc(allocator);

    const path = try std.fs.path.join(allocator, &[_][]const u8{cwd, day});
    std.os.mkdir(path, 511) catch |err| {
        if (err != std.os.MakeDirError.PathAlreadyExists) {
            return err;
        }
        std.log.info("folder '{s}' already exists, skipping creation", .{path});
    };

    const input_dir = try std.fs.path.join(allocator, &[_][]const u8{cwd, "../input/", day});
    std.os.mkdir(input_dir, 511) catch |err| {
        if (err != std.os.MakeDirError.PathAlreadyExists) {
            return err;
        }
        std.log.info("folder '{s}' already exists, skipping creation", .{input_dir});
    };

    const example_path = try std.fs.path.join(allocator, &[_][]const u8{input_dir, "example.txt"});
    _ = try std.fs.createFileAbsolute(example_path, .{.read = true});

    const question_path = try std.fs.path.join(allocator, &[_][]const u8{input_dir, "question.txt"});
    _ = try std.fs.createFileAbsolute(question_path, .{.read = true});

    const code_path = try std.fs.path.join(allocator, &[_][]const u8{path, "part1.zig"});
    _ = try std.fs.createFileAbsolute(code_path, .{.read = true});
}