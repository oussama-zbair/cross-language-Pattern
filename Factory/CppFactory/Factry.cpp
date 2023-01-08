#include <iostream>

// Abstract base class for factories
class AbstractFactory {
 public:
  virtual ~AbstractFactory() = default;
  virtual void createProductA() = 0;
  virtual void createProductB() = 0;
};

// Concrete factory for creating concrete products
class ConcreteFactory : public AbstractFactory {
 public:
  void createProductA() override {
    std::cout << "Creating ConcreteProductA!" << std::endl;
  }
  void createProductB() override {
    std::cout << "Creating ConcreteProductB!" << std::endl;
  }
};

int main() {
  ConcreteFactory factory;
  factory.createProductA();
  factory.createProductB();
  return 0;
}
