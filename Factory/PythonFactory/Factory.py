# Abstract base class for factories
class AbstractFactory:
    def create_product_a(self):
        raise NotImplementedError
    
    def create_product_b(self):
        raise NotImplementedError

# Concrete factory for creating concrete products
class ConcreteFactory(AbstractFactory):
    def create_product_a(self):
        return ConcreteProductA()
    
    def create_product_b(self):
        return ConcreteProductB()

# Abstract base class for products
class AbstractProductA:
    def do_something(self):
        raise NotImplementedError

class AbstractProductB:
    def do_something(self):
        raise NotImplementedError

# Concrete products
class ConcreteProductA(AbstractProductA):
    def do_something(self):
        print("ConcreteProductA is doing something!")

class ConcreteProductB(AbstractProductB):
    def do_something(self):
        print("ConcreteProductB is doing something!")

# Client code
factory = ConcreteFactory()
product_a = factory.create_product_a()
product_b = factory.create_product_b()
product_a.do_something()
product_b.do_something()
