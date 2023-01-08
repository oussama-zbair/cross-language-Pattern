# Define the product interface
class Product
    def use
      raise NotImplementedError
    end
  end
  
  # Define the concrete products
  class ConcreteProductA < Product
    def use
      puts "Using ConcreteProductA"
    end
  end
  
  class ConcreteProductB < Product
    def use
      puts "Using ConcreteProductB"
    end
  end
  
  # Define the factory interface
  class Factory
    def create_product
      raise NotImplementedError
    end
  end
  
  # Define the concrete factories
  class ConcreteFactoryA < Factory
    def create_product
      ConcreteProductA.new
    end
  end
  
  class ConcreteFactoryB < Factory
    def create_product
      ConcreteProductB.new
    end
  end
  
  # Client code
  factory = ConcreteFactoryA.new
  product = factory.create_product
  product.use
  
  factory = ConcreteFactoryB.new
  product = factory.create_product
  product.use
  