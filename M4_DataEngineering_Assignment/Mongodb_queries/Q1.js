db.users.aggregate([
    {
      $lookup: {
        from: "orders",
        localField: "userId",
        foreignField: "userId",
        as: "userOrders"
      }
    },
    {
      $unwind: "$userOrders"
    },
    {
      $group: {
        _id: "$userId",
        totalSpent: { $sum: "$userOrders.totalAmount" }
      }
    },
    {
      $match: {
        totalSpent: { $gt: 500 }
      }
    },
    {
      $lookup: {
        from: "users",
        localField: "_id",
        foreignField: "userId",
        as: "userDetails"
      }
    },
    {
      $project: {
        _id: 0,
        userId: "$_id",
        totalSpent: 1,
        name: { $arrayElemAt: ["$userDetails.name", 0] }
      }
    }
  ])
  