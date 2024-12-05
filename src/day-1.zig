const std = @import("std");

pub fn main() !void {
    // Read bytes
    const input = @embedFile("input/day-1.txt");

    // Parse into list
    var gpa = std.heap.GeneralPurposeAllocator(.{}).init;
    defer _ = gpa.deinit();
    const ally = gpa.allocator();

    var list_1 = std.ArrayList(i32).init(ally);
    var list_2 = std.ArrayList(i32).init(ally);
    defer list_1.deinit();
    defer list_2.deinit();

    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |line| {
        var columns = std.mem.splitSequence(u8, line, "   ");

        // EAFP approach instead of LBYL when reaching EOF
        const left = std.fmt.parseInt(i32, columns.first(), 0) catch break;
        const right = std.fmt.parseInt(i32, columns.rest(), 0) catch break;

        try list_1.append(left);
        try list_2.append(right);
    }

    // Ensure that EAFP works
    std.debug.assert(list_1.items.len == list_1.items.len);

    // Sort
    std.mem.sort(i32, list_1.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, list_2.items, {}, comptime std.sort.asc(i32));

    // Sum difference
    var sum: u32 = 0;
    var i: usize = 0;
    while (i < list_1.items.len) : (i += 1) {
        sum += @abs(list_1.items[i] - list_2.items[i]);
    }

    // Print
    const std_out = std.io.getStdOut().writer();
    try std_out.print("{}\n", .{sum});
}
