package org.example;

import javax.annotation.Nullable;
import java.util.*;


public class Canva {
    // Comment class
    public static class Comment {
        public final int commentId;
        public String ownerId;
        public String body;

        public Optional<Integer> parentId;

        public int points;

        public Comment(int commentId, String ownerId, String body, Integer parentId) {
            this.commentId = commentId;
            this.ownerId = ownerId;
            this.body = body;
            this.points = 0;
            if (parentId == null) {
                this.parentId = Optional.empty();
            } else {
                this.parentId = Optional.of(parentId);
            }
        }
    }

    // CommentService class

    /**
     * /comments?page=1
     * /comments/vote?=direction
     *
     * POST a comment
     * POST Reply to a comment
     * PUT Edit comments
     * DELETE Delete comments
     *
     * POST Upvote / Downvote comment
     * GET Get comments for a given post
     */
    public static class CommentService {
        private int idCounter = 0;
        private Map<Integer, Comment> store = new HashMap<>();

        /**
         * Post a new comment, or reply to comment
         * @param ownerId
         * @param body
         * @param parentId ID of parent comment. null if root comment.
         * @return commentId
         */
        public int postComment(String ownerId, String body, @Nullable Integer parentId) {
            int commentId = idCounter++;
            Comment comment = new Comment(commentId, ownerId, body, parentId);
            store.put(commentId, comment);
            return commentId;
        }

        public int editComment(String body, int commentId) {
            Comment comment = store.get(commentId);
            //if (ownerId.equals(store.get(commentId).ownerId))
            comment.body = body;
            store.put(commentId, comment);
            return commentId;
        }

        public void upvote(User user, int commentId) {
            if (user.commentsUpvoted.contains(commentId)) {
                throw new RuntimeException("invalid, 400");
            }
            store.get(commentId).points++;
        }

        public void delete(int commentId) {
            //if (ownerId.equals(store.get(commentId).ownerId))
            Comment comment = store.get(commentId);
            comment.body = "<DELETED>";
            comment.ownerId = "<DELETED>";
        }

        public List<Comment> getAll() {
            return new ArrayList<>(store.values());
        }
    }

    public static class User {
        public ArrayList<Integer> commentsUpvoted = new ArrayList<>();
        public ArrayList<Integer> commentsDownvoted = new ArrayList<>();
    }

}