class Squares
  def initialize(max)
    @max = max
  end

  def square_of_sums
    (1..@max).inject(:+)**2
  end

  def sum_of_squares
    (1..@max).inject { |sum, number| sum + number**2 }
  end

  def difference
    square_of_sums - sum_of_squares
  end
end