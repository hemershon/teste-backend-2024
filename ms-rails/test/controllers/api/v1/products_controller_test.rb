# test/controllers/api/v1/products_controller_test.rb
require 'test_helper'

class Api::V1::ProductsControllerTest < ActionController::TestCase
  setup do
    @product = Product.create(name: "Product", brand: "Brand", price: 10.0, description: "Description", stock: 100)
  end

  test "should get show" do
    assert_not_nil @product
    get :show, params: { id: @product.id }
    assert_response :success
  end

  test "should update product" do
    assert_not_nil @product
    put :update, params: { id: @product.id, product: { name: "Laptop", brand: "keyd" } }
    @product.reload
    assert_equal "Laptop", @product.name
    assert_response :success
  end

end
