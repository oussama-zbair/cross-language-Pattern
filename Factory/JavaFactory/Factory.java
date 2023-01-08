import java.util.HashMap;
import java.util.Map;

abstract class Product {
    // Abstract method that must be implemented by concrete products
    public abstract void use();
}

class ConcreteProductA extends Product {
    @Override
    public void use() {
        System.out.println("Using ConcreteProductA");
    }
}

class ConcreteProductB extends Product {
    @Override
    public void use() {
        System.out.println("Using ConcreteProductB");
    }
}

abstract class Creator {
    // Map to store all products
    private static Map<String, Product> products = new HashMap<>();

    // Factory method that returns an instance of a concrete product
    protected abstract Product factoryMethod(String productType);

    // Method to get a product instance by product type
    public Product createProduct(String productType) {
        Product product = products.get(productType);
        if (product == null) {
            product = factoryMethod(productType);
            products.put(productType, product);
        }
        return product;
    }
}

class ConcreteCreatorA extends Creator {
    @Override
    protected Product factoryMethod(String productType) {
        switch (productType) {
            case "A":
                return new ConcreteProductA();
            case "B":
                return new ConcreteProductB();
            default:
                throw new IllegalArgumentException("Invalid product type");
        }
    }
}

class Client {
    public static void main(String[] args) {
        Creator creator = new ConcreteCreatorA();

        Product productA = creator.createProduct("A");
        productA.use();

        Product productB = creator.createProduct("B");
        productB.use();
    }
}
