class Complement
  VERSION = 1
  DNA_TO_RNA = {
    "G" => "C",
    "C" => "G",
    "T" => "A",
    "A" => "U"
  }

  def self.of_dna(strand)
    of_strand(strand, :of_dna_base)
  end

  def self.of_rna(strand)
    of_strand(strand, :of_rna_base)
  end

  def self.of_strand(strand, of_base)
    strand.each_char.inject("") do |complement, base|
      complement + send(of_base, base)
    end
  end

  def self.of_dna_base(base)
    DNA_TO_RNA.fetch(base) { raise ArgumentError }
  end

  def self.of_rna_base(base)
    raise ArgumentError unless DNA_TO_RNA.value?(base)
    DNA_TO_RNA.key(base)
  end
end