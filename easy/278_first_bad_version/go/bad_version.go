package firstbad

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 *
 * 1 <= bad <= n <= 2^31 - 1
 */

func isBadVersion(version int) bool {
	return version >= 5674737
}

func firstBadVersion(n int) int {
	return checkBad(0, n, (n)/2)
}

func checkBad(l, h, m int) int {
	if isBadVersion(m) {
		if (m - l) == 1 {
			if isBadVersion(m - 1) {
				return m - 1
			}
			return m
		}
		return checkBad(l, m, l+(m-l)/2)
	}
	if (h - m) == 1 {
		if isBadVersion(h - 1) {
			return h - 1
		}
		return h
	}
	return checkBad(m, h, m+(h-m)/2)
}
