class CreateShippingAddresses < ActiveRecord::Migration
  def change
    create_table :shipping_addresses do |t|
      t.string :name
      t.string :address
      t.string :city
      t.string :zip
      t.string :state
      t.string :phone
      t.string :email
      t.references :shipping_method, index: true, foreign_key: true
      t.references :customer, index: true, foreign_key: true
      t.timestamps null: false
    end
  end
end
