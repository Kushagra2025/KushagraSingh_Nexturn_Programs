db.products.aggregate([
    {
      $unwind: "$ratings"
    },
    {
      $group: {
        _id: "$productId",
        averageRating: { $avg: "$ratings.rating" }
      }
    },
    {
      $match: {
        averageRating: { $gte: 4 }
      }
    },
    {
      $lookup: {
        from: "products",
        localField: "_id",
        foreignField: "productId",
        as: "productDetails"
      }
    },
    {
      $project: {
        _id: 0,
        productId: "$_id",
        averageRating: 1,
        name: { $arrayElemAt: ["$productDetails.name", 0] }
      }
    }
  ])
  