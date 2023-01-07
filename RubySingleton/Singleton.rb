class Singleton
    @@instance = nil
  
    def self.instance
      @@instance = new unless @@instance
      @@instance
    end
  
    private_class_method :new
  end
  