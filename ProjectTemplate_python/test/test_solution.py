from solution.solution import solution
import unittest
import pytest


class TestSolution(unittest.TestCase):

    @pytest.mark.timeout(1.00)
    def test_solution(self):
        r = solution([1, 2, 3])
        self.assertEquals(-1, r)
