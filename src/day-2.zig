const std = @import("std");
const math = @import("std").math;

const input: []const u8 = @embedFile("input/day-2.txt");

pub fn main() !void {
    var total_safe: i32 = 0;

    var reports = std.mem.tokenizeScalar(u8, input, '\n');

    while (reports.next()) |report| {
        var levels = std.mem.tokenizeScalar(u8, report, ' ');

        const first = try std.fmt.parseInt(i32, levels.next().?, 0);
        const second = try std.fmt.parseInt(i32, levels.next().?, 0);

        if (first == second) continue;
        const diff = @abs(first - second);
        if (diff < 1 or diff > 3) continue;

        const is_asc = first < second;
        var previous = second;
        var ok = true;

        while (levels.next()) |level| {
            const current = try std.fmt.parseInt(i32, level, 0);

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

    const std_out = std.io.getStdOut().writer();
    try std_out.print("{}", .{total_safe});
}
