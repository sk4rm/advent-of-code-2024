const std = @import("std");
const math = @import("std").math;

const input: []const u8 = @embedFile("input/day-2.txt");
var problem_dampener_enabled = false;

pub fn main() !void {
    // Parse input

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

    // Solution driver

    const total_safe: i32 = solve(reports);

    const std_out = std.io.getStdOut().writer();
    try std_out.print("{}", .{total_safe});
}

fn solve(reports: std.ArrayList(std.ArrayList(i32))) i32 {
    var total_safe: i32 = 0;

    for (reports.items) |report| {
        const ok = checkReport(report, problem_dampener_enabled);
        if (ok) total_safe += 1;
    }

    return total_safe;
}

fn checkReport(report: std.ArrayList(i32), can_remove_level: bool) bool {
    const first_level = report.items[0];
    const second_level = report.items[1];

    if (first_level == second_level) return false;
    const diff = @abs(first_level - second_level);
    if (diff < 1 or diff > 3) return false;

    const is_asc = first_level < second_level;
    var previous = second_level;
    var ok = true;

    for (report.items[2..]) |current| {
        if (previous == current) {
            ok = false;
        } else if (is_asc) {
            ok = (current - previous >= 1) and (current - previous <= 3);
        } else {
            ok = (previous - current >= 1) and (previous - current <= 3);
        }

        if (!ok) {
            if (!can_remove_level) return false;

            // Part 2: test all combinations of removing 1 level from report
        }

        previous = current;
    }

    // return false?
    return ok;
}
