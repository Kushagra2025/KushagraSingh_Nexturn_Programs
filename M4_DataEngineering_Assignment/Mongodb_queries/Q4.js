db.orders.find({ orderId: "ORD001" }).forEach(function(order) {
    order.items.forEach(function(item) {
      db.products.updateOne(
        { productId: item.productId },
        { $inc: { stock: -item.quantity } }
      );
    });
  });
  