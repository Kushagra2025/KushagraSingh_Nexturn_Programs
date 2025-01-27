db.orders.aggregate([
    {
      $match: {
        orderDate: { 
          $gte: ISODate("2024-12-01T00:00:00Z"), 
          $lte: ISODate("2024-12-31T23:59:59Z")
        }
      }
    },
    {
      $lookup: {
        from: "users",
        localField: "userId",
        foreignField: "userId",
        as: "userDetails"
      }
    },
    {
      $unwind: "$userDetails"
    },
    {
      $project: {
        _id: 0,
        orderId: 1,
        orderDate: 1,
        userId: 1,
        userName: "$userDetails.name",
        totalAmount: 1
      }
    }
  ])
  