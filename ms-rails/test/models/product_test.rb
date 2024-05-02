# test/models/product_test.rb
require 'test_helper'

class ProductTest < ActiveSupport::TestCase
  test "should not save product without name" do
    product = Product.new
    assert_not product.save, "Saved the product without a name"
  end

  test "should not save product with short name" do
    product = Product.new(name: "abc")
    assert_not product.save, "Saved the product with a short name"
  end

  test "should not save product with invalid price format" do
    product = Product.new(name: "Product", price: "abc")
    assert_not product.save, "Saved the product with an invalid price format"
  end

  test "should not save product with price less than or equal to zero" do
    product = Product.new(name: "Product", price: 0)
    assert_not product.save, "Saved the product with price less than or equal to zero"
  end

  test "should not save product with price greater than or equal to 1000000" do
    product = Product.new(name: "Product", price: 1000000)
    assert_not product.save, "Saved the product with price greater than or equal to 1000000"
  end

  test "should not save product without brand" do
    product = Product.new(name: "Product", brand: nil)
    assert_not product.save, "Saved the product without a brand"
  end

  test "should not save product without description" do
    product = Product.new(name: "Product", description: nil)
    assert_not product.save, "Saved the product without a description"
  end

  test "should not save product without stock" do
    product = Product.new(name: "Product", stock: nil)
    assert_not product.save, "Saved the product without stock"
  end

  test "should not save product with non-integer stock" do
    product = Product.new(name: "Product", stock: 1.5)
    assert_not product.save, "Saved the product with non-integer stock"
  end
end
