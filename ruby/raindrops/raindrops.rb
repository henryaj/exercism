require 'prime'

class Raindrops
  DROPSPEAK = { 3 => "Pling", 5 => "Plang", 7 => "Plong" }

  def self.convert(number)
    factors = prime_factors(number)
    dropspeak = factors.map { |factor| DROPSPEAK[factor] }.join
    dropspeak.empty? ? number.to_s : dropspeak
  end

  def self.prime_factors(number)
    Prime.prime_division(number).collect(&:first)
  end
end