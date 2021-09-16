class Pilha

    def initialize
        @tamanho = 0
        @pilha = []
    end

    def tamanho
        @tamanho
    end
    
    def empilhar(obj)
        @pilha << obj
        @tamanho += 1
    end

    def desempilhar
        @tamanho -=1
        @pilha.pop
    end
end

pilha = Pilha.new

pilha.empilhar(1)
pilha.empilhar(2)
pilha.empilhar(3)
pilha.empilhar(4)

puts pilha.tamanho

puts pilha.desempilhar
puts pilha.desempilhar
puts pilha.desempilhar
puts pilha.desempilhar

puts pilha.tamanho
