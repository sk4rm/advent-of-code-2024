const std = @import("std");

pub fn main() !void {
    const input = @embedFile("input/day-1.txt");
    std.debug.print("{s}", .{input});
}
