/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        return dfs(cloned, target);
    }

    TreeNode* dfs(TreeNode* root, TreeNode* target) {
        if (root == NULL) {
            return NULL;
        }

        if (root->val == target->val) {
            return root;
        }

        auto left = dfs(root->left, target);
        if (left != NULL) {
            return left;
        }
        auto right = dfs(root->right, target);
        if (right != NULL) {
            return right;
        }

        return NULL;
    }
};