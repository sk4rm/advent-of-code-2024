const std = @import("std");
const math = @import("std").math;

const input: []const u8 = @embedFile("input/day-2.txt");
var problem_dampener_enabled = true;

pub fn main() !void {
    // Memory allocation and management

    var gpa = std.heap.GeneralPurposeAllocator(.{}).init;
    var arena = std.heap.ArenaAllocator.init(gpa.allocator());
    defer arena.deinit();
    const ally = arena.allocator();

    var reports = std.ArrayList(std.ArrayList(i32)).init(ally);
    defer reports.deinit();

    var lines = std.mem.tokenizeScalar(u8, input, '\n');
    while (lines.next()) |line| {
        var report = std.ArrayList(i32).init(ally);

        var levels = std.mem.tokenizeScalar(u8, line, ' ');
        while (levels.next()) |level| {
            try report.append(try std.fmt.parseInt(i32, level, 0));
        }

        try reports.append(report);
    }

    // Actual solving

    const total_safe: i32 = solve(reports);

    const std_out = std.io.getStdOut().writer();
    try std_out.print("{}", .{total_safe});
}

fn solve(reports: std.ArrayList(std.ArrayList(i32))) i32 {
    var total_safe: i32 = 0;
    var i: i32 = 1;

    for (reports.items) |levels| {
        defer i += 1;
        const first = levels.items[0];
        const second = levels.items[1];

        if (first == second) continue;
        const diff = @abs(first - second);
        if (diff < 1 or diff > 3) continue;

        const is_asc = first < second;
        var previous = second;
        var ok = true;

        for (levels.items[2..]) |current| {
            if (previous == current) {
                ok = false;
            } else if (is_asc) {
                ok = (current - previous >= 1) and (current - previous <= 3);
            } else {
                ok = (previous - current >= 1) and (previous - current <= 3);
            }

            if (!ok) break;
            previous = current;
        }

        if (ok) total_safe += 1;
    }

    return total_safe;
}
